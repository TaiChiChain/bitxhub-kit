package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/axiomesh/axiom-kit/hexutil"
	"github.com/axiomesh/axiom-kit/types/pb"
)

type Node interface {
	Encode() []byte
	GetHash() common.Hash
	Copy() Node
	Type() int
	String() string // just for debug
}

type (
	TrieJournal struct {
		Type        byte
		DirtySet    map[string]Node
		PruneSet    map[string]struct{}
		RootHash    common.Hash
		RootNodeKey *NodeKey
	}

	SnapJournal struct {
		Destruct map[string]struct{}
		Account  map[string][]byte
		Storage  map[string]map[string][]byte
	}

	StateJournal struct {
		RootHash        *Hash
		CodeJournal     map[string][]byte
		TrieJournal     []*TrieJournal
		SnapshotJournal *SnapJournal
	}
)

type (
	InternalNode struct {
		Children [TrieDegree]*Child

		blob []byte      // encode result for reuse
		hash common.Hash // hash result for reuse
	}

	LeafNode struct {
		Key  []byte
		Val  []byte
		Hash common.Hash
	}

	Child struct {
		Hash    common.Hash
		Version uint64
		Leaf    bool
	}
)

type (
	NodeKey struct {
		Version uint64 // version of current tree node.
		Path    []byte // addressing path from root to current node. Path is part of LeafNode.Key.
		Type    []byte // additional field for identify a tree uniquely together with Version and Path.
	}

	NodeKeyHeap []*NodeKey // max heap to store NodeKey
)

const (
	TrieDegree       = 16
	TypeInternalNode = 0
	TypeLeafNode     = 1
)

var internalNodePool = sync.Pool{
	New: func() any {
		return &InternalNode{}
	},
}

var diffPool = sync.Pool{
	New: func() any {
		return &StateJournal{}
	},
}

var childPool = sync.Pool{
	New: func() any {
		return &Child{}
	},
}

var (
	ErrorEmptyStateJournal = errors.New("decode empty state journal")
)

// just for debug
func (nk *NodeKey) String() string {
	res := strings.Builder{}
	res.WriteString("Version[")
	res.WriteString(strconv.Itoa(int(nk.Version)))
	res.WriteString("], Type[")
	res.WriteString(strconv.Itoa(len(nk.Type)))
	res.WriteString("], Path[")
	res.WriteString(hexutil.DecodeFromNibbles(nk.Path))
	res.WriteString("]")
	return res.String()
}

func (nk *NodeKey) Encode() []byte {
	var length byte
	for i := 0; i < len(nk.Type); i++ {
		length++
	}
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, nk.Version)
	buf = append(buf, length)
	buf = append(buf, nk.Type...)
	buf = append(buf, nk.Path...)
	return buf
}

func (nk *NodeKey) GetIndex() []byte {
	var length byte
	for i := 0; i < len(nk.Type); i++ {
		length++
	}
	buf := make([]byte, 0)
	buf = append(buf, length)
	buf = append(buf, nk.Type...)
	buf = append(buf, nk.Path...)
	return buf
}

// DecodeNodeKey decode from bytes in physical storage to NodeKey
func DecodeNodeKey(raw []byte) *NodeKey {
	nk := &NodeKey{}
	nk.Version = binary.BigEndian.Uint64(raw[:8])
	length := raw[8]
	nk.Type = raw[9 : 9+length]
	nk.Path = raw[9+length:]
	return nk
}

func RecycleTrieNode(n Node) {
	if n != nil && n.Type() == TypeInternalNode {
		nn := n.(*InternalNode)
		nn.blob = nil
		nn.hash = common.Hash{}
		for i := range nn.Children {
			if nn.Children[i] == nil {
				continue
			}
			nn.Children[i].Hash = common.Hash{}
			nn.Children[i].Leaf = false
			nn.Children[i].Version = 0
			childPool.Put(nn.Children[i])
			nn.Children[i] = nil
		}
		internalNodePool.Put(nn)
	}
}

func RecycleStateJournal(diff *StateJournal) {
	if diff == nil {
		return
	}
	for _, trieJournal := range diff.TrieJournal {
		for _, node := range trieJournal.DirtySet {
			RecycleTrieNode(node)
		}
	}
	diff.RootHash = nil
	diff.SnapshotJournal = nil
	diff.CodeJournal = nil
	diff.TrieJournal = nil

	diffPool.Put(diff)
}

// just for debug
func (n *InternalNode) String() string {
	res := strings.Builder{}
	for i := 0; i < TrieDegree; i++ {
		if n.Children[i] == nil {
			continue
		} else {
			res.WriteString(strconv.Itoa(i))
			child := n.Children[i]
			res.WriteString(":<Hash[")
			res.WriteString(child.Hash.String()[2:6])
			res.WriteString("], Version[")
			res.WriteString(strconv.Itoa(int(child.Version)))
			res.WriteString("], Leaf[")
			res.WriteString(strconv.FormatBool(child.Leaf))
			res.WriteString("]>, ")
		}
	}
	return res.String()
}

// just for debug
func (n *LeafNode) String() string {
	res := strings.Builder{}
	res.WriteString("Key[")
	res.WriteString(hexutil.DecodeFromNibbles(n.Key))
	res.WriteString("], ValueLen[")
	res.WriteString(strconv.Itoa(len(n.Val)))
	res.WriteString("], Hash[")
	res.WriteString(n.Hash.String()[2:6])
	res.WriteString("]")
	return res.String()
}

// just for debug
func (j *TrieJournal) String() string {
	res := strings.Builder{}
	res.WriteString("Version[")
	res.WriteString(strconv.Itoa(int(j.RootNodeKey.Version)))
	res.WriteString("], ")

	res.WriteString(fmt.Sprintf("{DirtySet[\n"))
	dirtyKeys := make([]string, 0, len(j.DirtySet))
	for k := range j.DirtySet {
		dirtyKeys = append(dirtyKeys, k)
	}
	sort.Strings(dirtyKeys)
	for _, k := range dirtyKeys {
		// res.WriteString(fmt.Sprintf("%v\n", DecodeNodeKey([]byte(k)).String()))
		res.WriteString(fmt.Sprintf("%v\n", []byte(k)))
	}

	res.WriteString("], PruneSet[\n")
	pruneKeys := make([]string, 0, len(j.PruneSet))
	for k := range j.PruneSet {
		pruneKeys = append(pruneKeys, k)
	}
	sort.Strings(pruneKeys)
	for _, k := range pruneKeys {
		// res.WriteString(fmt.Sprintf("%v\n", DecodeNodeKey([]byte(k)).String()))
		res.WriteString(fmt.Sprintf("%v\n", []byte(k)))
	}
	res.WriteString("]}\n")

	return res.String()
}

func (n *InternalNode) GetHash() common.Hash {
	if n.hash != (common.Hash{}) {
		return n.hash
	}
	data := sha256.Sum256(n.Encode())
	n.hash = data
	return data
}

// GetHash get LeafNode's hash, which is only determined by its Key and Val
func (n *LeafNode) GetHash() common.Hash {
	tmp := &LeafNode{
		Hash: common.Hash{},
		Key:  n.Key,
		Val:  n.Val,
	}
	data := sha256.Sum256(tmp.Encode())
	return data
}

// deep copy
func (n *InternalNode) Copy() Node {
	nn := internalNodePool.Get().(*InternalNode)
	for i, child := range n.Children {
		if child == nil {
			continue
		}
		nn.Children[i] = &Child{
			Hash:    child.Hash,
			Leaf:    child.Leaf,
			Version: child.Version,
		}
	}
	return nn
}

// deep copy
func (n *LeafNode) Copy() Node {
	nn := &LeafNode{
		Hash: n.Hash,
		Key:  make([]byte, len(n.Key)),
		Val:  make([]byte, len(n.Val)),
	}
	copy(nn.Key, n.Key)
	copy(nn.Val, n.Val)
	return nn
}

func (n *InternalNode) Type() int {
	return TypeInternalNode
}

func (n *LeafNode) Type() int {
	return TypeLeafNode
}

func (n *InternalNode) Encode() []byte {
	if n == nil {
		return nil
	}

	if len(n.blob) > 0 {
		return n.blob
	}

	children := make([]*pb.Child, TrieDegree)
	for i, child := range n.Children {
		if child == nil {
			continue
		}
		children[i] = &pb.Child{
			Version: child.Version,
			Leaf:    child.Leaf,
			Hash:    child.Hash[:],
		}
	}

	blob := &pb.InternalNode{
		Children: children,
	}

	content, err := blob.MarshalVTStrict()
	if err != nil {
		return nil
	}

	node := &pb.Node{
		Leaf:    false,
		Content: content,
	}
	res, err := node.MarshalVTStrict()
	if err != nil {
		return nil
	}
	n.blob = res
	return res
}

func (n *InternalNode) unmarshalInternalFromPb(data []byte) error {
	helper := pb.InternalNodeFromVTPool()
	defer func() {
		helper.Reset()
		helper.ReturnToVTPool()
	}()
	err := helper.UnmarshalVT(data)
	if err != nil {
		return err
	}

	n.Children = [TrieDegree]*Child{}
	for i, child := range helper.Children {
		if len(child.Hash) == 0 {
			continue
		}
		n.Children[i] = childPool.Get().(*Child)
		n.Children[i].Hash = common.BytesToHash(child.Hash)
		n.Children[i].Version = child.Version
		n.Children[i].Leaf = child.Leaf
	}
	return nil
}

func (n *LeafNode) Encode() []byte {
	if n == nil {
		return nil
	}

	blob := &pb.LeafNode{
		Key:   HexToBytes(n.Key),
		Value: n.Val,
		Hash:  n.Hash[:],
	}

	content, err := blob.MarshalVTStrict()
	if err != nil {
		return nil
	}

	node := &pb.Node{
		Leaf:    true,
		Content: content,
	}
	res, err := node.MarshalVTStrict()
	if err != nil {
		return nil
	}
	return res
}

func (n *LeafNode) unmarshalLeafFromPb(data []byte) error {
	helper := pb.LeafNodeFromVTPool()
	defer func() {
		helper.Reset()
		helper.ReturnToVTPool()
	}()
	err := helper.UnmarshalVT(data)
	if err != nil {
		return err
	}
	n.Val = helper.Value
	n.Key = BytesToHex(helper.Key)
	n.Hash = common.BytesToHash(helper.Hash)
	return nil
}

func UnmarshalJMTNodeFromPb(data []byte) (Node, error) {
	if len(data) == 0 {
		return nil, nil
	}
	helper := pb.NodeFromVTPool()
	defer func() {
		helper.Reset()
		helper.ReturnToVTPool()
	}()
	err := helper.UnmarshalVT(data)
	if err != nil {
		return nil, err
	}

	if helper.Leaf {
		res := &LeafNode{}
		err = res.unmarshalLeafFromPb(helper.Content)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	if !helper.Leaf {
		res := internalNodePool.Get().(*InternalNode)
		err = res.unmarshalInternalFromPb(helper.Content)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	return nil, nil
}

// BytesToHex expand normal bytes to hex bytes (nibbles)
func BytesToHex(h []byte) []byte {
	if len(h) == 0 {
		return h
	}
	var i byte = 0
	length := len(h) * 2
	dst := make([]byte, length)
	for ; int(i) < length; i++ {
		if i&0x01 == 0 {
			dst[i] = h[i/2] >> 4
		} else {
			dst[i] = h[i/2] & 15
		}
	}
	return dst
}

// HexToBytes compress hex bytes (also called nibbles) to normal bytes
func HexToBytes(src []byte) []byte {
	if len(src) == 0 {
		return src
	}
	if len(src) > 255 {
		panic("don't support compress bytes with length greater than 255 ")
	}
	if len(src)%2 != 0 {
		panic("don't support compress bytes with odd length")
	}
	length := len(src) / 2
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[i] = (src[2*i] << 4) | src[2*i+1]
	}
	return res
}

func (diff *StateJournal) Encode() []byte {
	if diff == nil {
		return nil
	}

	journals := make([]*pb.TrieJournal, len(diff.TrieJournal))
	for i, journal := range diff.TrieJournal {
		pruneSet := make(map[string][]byte)
		for k := range journal.PruneSet {
			pruneSet[k] = []byte{}
		}

		dirtySet := make(map[string][]byte)
		for k, v := range journal.DirtySet {
			dirtySet[k] = v.Encode()
		}

		journals[i] = &pb.TrieJournal{
			Type:        uint64(journal.Type),
			RootHash:    journal.RootHash[:],
			RootNodeKey: journal.RootNodeKey.Encode(),
			PruneSet:    pruneSet,
			DirtySet:    dirtySet,
		}
	}

	snapAccount := make(map[string][]byte)
	snapStorage := make(map[string]*pb.SnapStorage)
	snapDestruct := make(map[string][]byte)
	for k, v := range diff.SnapshotJournal.Account {
		snapAccount[k] = v
	}
	for k, v := range diff.SnapshotJournal.Storage {
		snapStorage[k] = &pb.SnapStorage{SnapStorage: v}
	}
	for k := range diff.SnapshotJournal.Destruct {
		snapDestruct[k] = []byte{}
	}

	snapJournal := &pb.SnapJournal{
		Account:  snapAccount,
		Storage:  snapStorage,
		Destruct: snapDestruct,
	}

	blob := &pb.StateJournal{
		RootHash:    diff.RootHash.Bytes(),
		TrieJournal: journals,
		CodeJournal: diff.CodeJournal,
		SnapJournal: snapJournal,
	}

	content, err := blob.MarshalVTStrict()
	if err != nil {
		return nil
	}

	return content
}

func DecodeStateJournal(data []byte) (*StateJournal, error) {
	if len(data) == 0 {
		return nil, ErrorEmptyStateJournal
	}
	helper := pb.StateJournalFromVTPool()
	defer func() {
		helper.Reset()
		helper.ReturnToVTPool()
	}()
	err := helper.UnmarshalVT(data)
	if err != nil {
		return nil, err
	}

	diff := diffPool.Get().(*StateJournal)
	diff.TrieJournal = make([]*TrieJournal, len(helper.TrieJournal))
	diff.CodeJournal = helper.CodeJournal
	diff.SnapshotJournal = &SnapJournal{
		Account:  helper.SnapJournal.Account,
		Storage:  make(map[string]map[string][]byte),
		Destruct: make(map[string]struct{}),
	}

	for i := 0; i < len(helper.TrieJournal); i++ {
		pruneSet := make(map[string]struct{})
		for k := range helper.TrieJournal[i].PruneSet {
			pruneSet[k] = struct{}{}
		}

		dirtySet := make(map[string]Node)
		for k, v := range helper.TrieJournal[i].DirtySet {
			dirtySet[k], err = UnmarshalJMTNodeFromPb(v)
			if err != nil {
				return nil, err
			}
		}

		diff.TrieJournal[i] = &TrieJournal{
			Type:        byte(helper.TrieJournal[i].Type),
			RootHash:    common.BytesToHash(helper.TrieJournal[i].RootHash),
			RootNodeKey: DecodeNodeKey(helper.TrieJournal[i].RootNodeKey),
			PruneSet:    pruneSet,
			DirtySet:    dirtySet,
		}
	}

	for k := range helper.SnapJournal.Destruct {
		diff.SnapshotJournal.Destruct[k] = struct{}{}
	}
	for k, v := range helper.SnapJournal.Storage {
		diff.SnapshotJournal.Storage[k] = v.SnapStorage
	}

	diff.RootHash = NewHash(helper.RootHash)

	return diff, nil
}

func (h NodeKeyHeap) Len() int {
	return len(h)
}

func (h NodeKeyHeap) Less(i, j int) bool {
	if h[i].Version != h[j].Version {
		return h[i].Version > h[j].Version
	}

	if !bytes.Equal(h[i].Type, h[j].Type) {
		return bytes.Compare(h[i].Type, h[j].Type) > 0
	}

	return bytes.Compare(h[i].Path, h[j].Path) < 0
}

func (h NodeKeyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeKeyHeap) Push(x any) {
	*h = append(*h, x.(*NodeKey))
}

func (h *NodeKeyHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
