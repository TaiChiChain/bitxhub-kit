package types

import (
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/axiomesh/axiom-kit/hexutil"
	"github.com/axiomesh/axiom-kit/types/pb"
)

type Node interface {
	EncodePb() []byte
	EncodeJson() []byte
	GetHash() common.Hash
	Copy() Node
	String() string // just for debug
}

type (
	InternalNode struct {
		Children [16]*Child `json:"children"`
	}

	LeafNode struct {
		Key  []byte      `json:"key"`
		Val  []byte      `json:"value"`
		Hash common.Hash `json:"hash"`
	}

	Child struct {
		Hash    common.Hash `json:"hash"`
		Version uint64      `json:"version"`
		Leaf    bool        `json:"leaf"`
	}
)

func (n *InternalNode) EncodeJson() []byte {
	blob, err := json.Marshal(n)
	if err != nil {
		return nil
	}
	return blob
}

func (n *LeafNode) EncodeJson() []byte {
	blob, err := json.Marshal(n)
	if err != nil {
		return nil
	}
	return blob
}

// just for debug
func (n *InternalNode) String() string {
	res := strings.Builder{}
	for i := 0; i < 16; i++ {
		res.WriteString(strconv.Itoa(i))
		if n.Children[i] == nil {
			res.WriteString(", ")
		} else {
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

func (n *InternalNode) GetHash() common.Hash {
	data := sha256.Sum256(n.EncodePb())
	return data
}

// GetHash get LeafNode's hash, which is only determined by its Key and Val
func (n *LeafNode) GetHash() common.Hash {
	tmp := &LeafNode{
		Hash: common.Hash{},
		Key:  n.Key,
		Val:  n.Val,
	}
	data := sha256.Sum256(tmp.EncodePb())
	return data
}

func (n *InternalNode) Copy() Node {
	return &InternalNode{
		Children: n.Children,
	}
}

func (n *LeafNode) Copy() Node {
	nn := &LeafNode{
		Hash: n.Hash,
	}
	copy(nn.Key, n.Key)
	copy(nn.Val, n.Val)
	return nn
}

func (n *InternalNode) EncodePb() []byte {
	if n == nil {
		return nil
	}

	children := make([]*pb.Child, 16)
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

	n.Children = [16]*Child{}
	for i, child := range helper.Children {
		if len(child.Hash) == 0 {
			continue
		}
		n.Children[i] = &Child{
			Hash:    common.BytesToHash(child.Hash),
			Version: child.Version,
			Leaf:    child.Leaf,
		}
	}
	return nil
}

func (n *LeafNode) EncodePb() []byte {
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
		res := &InternalNode{}
		err = res.unmarshalInternalFromPb(helper.Content)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	return nil, nil
}

func UnmarshalJMTNodeFromJson(rawNode []byte) (n Node, err error) {
	if len(rawNode) == 0 {
		return nil, nil
	}
	m := make(map[string]any)
	if err := json.Unmarshal(rawNode, &m); err != nil {
		return nil, err
	}
	if m["children"] != nil {
		return unmarshalInternalFromJson(rawNode)
	} else if m["key"] != nil {
		return unmarshalLeafFromJson(rawNode)
	}
	return nil, nil
}

func unmarshalInternalFromJson(rawNode []byte) (Node, error) {
	n := &InternalNode{}
	err := json.Unmarshal(rawNode, n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func unmarshalLeafFromJson(rawNode []byte) (Node, error) {
	n := &LeafNode{}
	err := json.Unmarshal(rawNode, n)
	if err != nil {
		return nil, err
	}
	return n, nil
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
