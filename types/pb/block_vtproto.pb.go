// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.5.0
// source: block.proto

package pb

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (this *Block) EqualVT(that *Block) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.BlockHeader.EqualVT(that.BlockHeader) {
		return false
	}
	if len(this.Transactions) != len(that.Transactions) {
		return false
	}
	for i, vx := range this.Transactions {
		vy := that.Transactions[i]
		if string(vx) != string(vy) {
			return false
		}
	}
	if string(this.Extra) != string(that.Extra) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Block) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Block)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *BlockHeader) EqualVT(that *BlockHeader) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Number != that.Number {
		return false
	}
	if string(this.StateRoot) != string(that.StateRoot) {
		return false
	}
	if string(this.TxRoot) != string(that.TxRoot) {
		return false
	}
	if string(this.ReceiptRoot) != string(that.ReceiptRoot) {
		return false
	}
	if string(this.ParentHash) != string(that.ParentHash) {
		return false
	}
	if this.Timestamp != that.Timestamp {
		return false
	}
	if this.Epoch != that.Epoch {
		return false
	}
	if string(this.Bloom) != string(that.Bloom) {
		return false
	}
	if this.GasPrice != that.GasPrice {
		return false
	}
	if this.ProposerAccount != that.ProposerAccount {
		return false
	}
	if this.GasUsed != that.GasUsed {
		return false
	}
	if this.ProposerNodeId != that.ProposerNodeId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *BlockHeader) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*BlockHeader)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *Block) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Block) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Extra) > 0 {
		i -= len(m.Extra)
		copy(dAtA[i:], m.Extra)
		i = encodeVarint(dAtA, i, uint64(len(m.Extra)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Transactions) > 0 {
		for iNdEx := len(m.Transactions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transactions[iNdEx])
			copy(dAtA[i:], m.Transactions[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.Transactions[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockHeader != nil {
		size, err := m.BlockHeader.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeader) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *BlockHeader) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.ProposerNodeId != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ProposerNodeId))
		i--
		dAtA[i] = 0x68
	}
	if m.GasUsed != 0 {
		i = encodeVarint(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x60
	}
	if len(m.ProposerAccount) > 0 {
		i -= len(m.ProposerAccount)
		copy(dAtA[i:], m.ProposerAccount)
		i = encodeVarint(dAtA, i, uint64(len(m.ProposerAccount)))
		i--
		dAtA[i] = 0x5a
	}
	if m.GasPrice != 0 {
		i = encodeVarint(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Bloom) > 0 {
		i -= len(m.Bloom)
		copy(dAtA[i:], m.Bloom)
		i = encodeVarint(dAtA, i, uint64(len(m.Bloom)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Epoch != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x40
	}
	if m.Timestamp != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ParentHash) > 0 {
		i -= len(m.ParentHash)
		copy(dAtA[i:], m.ParentHash)
		i = encodeVarint(dAtA, i, uint64(len(m.ParentHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ReceiptRoot) > 0 {
		i -= len(m.ReceiptRoot)
		copy(dAtA[i:], m.ReceiptRoot)
		i = encodeVarint(dAtA, i, uint64(len(m.ReceiptRoot)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TxRoot) > 0 {
		i -= len(m.TxRoot)
		copy(dAtA[i:], m.TxRoot)
		i = encodeVarint(dAtA, i, uint64(len(m.TxRoot)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StateRoot) > 0 {
		i -= len(m.StateRoot)
		copy(dAtA[i:], m.StateRoot)
		i = encodeVarint(dAtA, i, uint64(len(m.StateRoot)))
		i--
		dAtA[i] = 0x12
	}
	if m.Number != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Block) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Block) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Extra) > 0 {
		i -= len(m.Extra)
		copy(dAtA[i:], m.Extra)
		i = encodeVarint(dAtA, i, uint64(len(m.Extra)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Transactions) > 0 {
		for iNdEx := len(m.Transactions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transactions[iNdEx])
			copy(dAtA[i:], m.Transactions[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.Transactions[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockHeader != nil {
		size, err := m.BlockHeader.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeader) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *BlockHeader) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.ProposerNodeId != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ProposerNodeId))
		i--
		dAtA[i] = 0x68
	}
	if m.GasUsed != 0 {
		i = encodeVarint(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x60
	}
	if len(m.ProposerAccount) > 0 {
		i -= len(m.ProposerAccount)
		copy(dAtA[i:], m.ProposerAccount)
		i = encodeVarint(dAtA, i, uint64(len(m.ProposerAccount)))
		i--
		dAtA[i] = 0x5a
	}
	if m.GasPrice != 0 {
		i = encodeVarint(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Bloom) > 0 {
		i -= len(m.Bloom)
		copy(dAtA[i:], m.Bloom)
		i = encodeVarint(dAtA, i, uint64(len(m.Bloom)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Epoch != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x40
	}
	if m.Timestamp != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ParentHash) > 0 {
		i -= len(m.ParentHash)
		copy(dAtA[i:], m.ParentHash)
		i = encodeVarint(dAtA, i, uint64(len(m.ParentHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ReceiptRoot) > 0 {
		i -= len(m.ReceiptRoot)
		copy(dAtA[i:], m.ReceiptRoot)
		i = encodeVarint(dAtA, i, uint64(len(m.ReceiptRoot)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TxRoot) > 0 {
		i -= len(m.TxRoot)
		copy(dAtA[i:], m.TxRoot)
		i = encodeVarint(dAtA, i, uint64(len(m.TxRoot)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StateRoot) > 0 {
		i -= len(m.StateRoot)
		copy(dAtA[i:], m.StateRoot)
		i = encodeVarint(dAtA, i, uint64(len(m.StateRoot)))
		i--
		dAtA[i] = 0x12
	}
	if m.Number != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

var vtprotoPool_Block = sync.Pool{
	New: func() interface{} {
		return &Block{}
	},
}

func (m *Block) ResetVT() {
	m.BlockHeader.ReturnToVTPool()
	f0 := m.Transactions[:0]
	f1 := m.Extra[:0]
	m.Reset()
	m.Transactions = f0
	m.Extra = f1
}
func (m *Block) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_Block.Put(m)
	}
}
func BlockFromVTPool() *Block {
	return vtprotoPool_Block.Get().(*Block)
}

var vtprotoPool_BlockHeader = sync.Pool{
	New: func() interface{} {
		return &BlockHeader{}
	},
}

func (m *BlockHeader) ResetVT() {
	f0 := m.StateRoot[:0]
	f1 := m.TxRoot[:0]
	f2 := m.ReceiptRoot[:0]
	f3 := m.ParentHash[:0]
	f4 := m.Bloom[:0]
	m.Reset()
	m.StateRoot = f0
	m.TxRoot = f1
	m.ReceiptRoot = f2
	m.ParentHash = f3
	m.Bloom = f4
}
func (m *BlockHeader) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_BlockHeader.Put(m)
	}
}
func BlockHeaderFromVTPool() *BlockHeader {
	return vtprotoPool_BlockHeader.Get().(*BlockHeader)
}
func (m *Block) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeader != nil {
		l = m.BlockHeader.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if len(m.Transactions) > 0 {
		for _, b := range m.Transactions {
			l = len(b)
			n += 1 + l + sov(uint64(l))
		}
	}
	l = len(m.Extra)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *BlockHeader) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sov(uint64(m.Number))
	}
	l = len(m.StateRoot)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.TxRoot)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.ReceiptRoot)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.ParentHash)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sov(uint64(m.Timestamp))
	}
	if m.Epoch != 0 {
		n += 1 + sov(uint64(m.Epoch))
	}
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.GasPrice != 0 {
		n += 1 + sov(uint64(m.GasPrice))
	}
	l = len(m.ProposerAccount)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.GasUsed != 0 {
		n += 1 + sov(uint64(m.GasUsed))
	}
	if m.ProposerNodeId != 0 {
		n += 1 + sov(uint64(m.ProposerNodeId))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Block) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockHeader == nil {
				m.BlockHeader = BlockHeaderFromVTPool()
			}
			if err := m.BlockHeader.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, make([]byte, postIndex-iNdEx))
			copy(m.Transactions[len(m.Transactions)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extra = append(m.Extra[:0], dAtA[iNdEx:postIndex]...)
			if m.Extra == nil {
				m.Extra = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockHeader) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRoot = append(m.StateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.StateRoot == nil {
				m.StateRoot = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxRoot = append(m.TxRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.TxRoot == nil {
				m.TxRoot = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiptRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiptRoot = append(m.ReceiptRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.ReceiptRoot == nil {
				m.ReceiptRoot = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentHash = append(m.ParentHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ParentHash == nil {
				m.ParentHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = append(m.Bloom[:0], dAtA[iNdEx:postIndex]...)
			if m.Bloom == nil {
				m.Bloom = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerNodeId", wireType)
			}
			m.ProposerNodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposerNodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
