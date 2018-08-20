package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	"github.com/CovenantSQL/HashStablePack/msgp"
)

// MarshalHash marshals for hash
func (z *Block) MarshalHash() (o []byte, err error) {
	var b []byte
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82, 0x82)
	if oTemp, err := z.SignedHeader.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = msgp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x82)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TxBillings)))
	for za0001 := range z.TxBillings {
		if z.TxBillings[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			if oTemp, err := z.TxBillings[za0001].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = msgp.AppendBytes(o, oTemp)
			}
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Block) Msgsize() (s int) {
	s = 1 + 13 + z.SignedHeader.Msgsize() + 11 + msgp.ArrayHeaderSize
	for za0001 := range z.TxBillings {
		if z.TxBillings[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.TxBillings[za0001].Msgsize()
		}
	}
	return
}

// MarshalHash marshals for hash
func (z *Header) MarshalHash() (o []byte, err error) {
	var b []byte
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	o = append(o, 0x85, 0x85)
	o = msgp.AppendInt32(o, z.Version)
	o = append(o, 0x85)
	if oTemp, err := z.Producer.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = msgp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x85)
	if oTemp, err := z.MerkleRoot.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = msgp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x85)
	if oTemp, err := z.ParentHash.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = msgp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x85)
	o = msgp.AppendTime(o, z.Timestamp)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Header) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int32Size + 9 + z.Producer.Msgsize() + 11 + z.MerkleRoot.Msgsize() + 11 + z.ParentHash.Msgsize() + 10 + msgp.TimeSize
	return
}

// MarshalHash marshals for hash
func (z *SignedHeader) MarshalHash() (o []byte, err error) {
	var b []byte
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	o = append(o, 0x84, 0x84)
	if oTemp, err := z.Header.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = msgp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x84)
	if oTemp, err := z.BlockHash.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = msgp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x84)
	if z.Signee == nil {
		o = msgp.AppendNil(o)
	} else {
		if oTemp, err := z.Signee.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = msgp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x84)
	if z.Signature == nil {
		o = msgp.AppendNil(o)
	} else {
		if oTemp, err := z.Signature.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = msgp.AppendBytes(o, oTemp)
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SignedHeader) Msgsize() (s int) {
	s = 1 + 7 + z.Header.Msgsize() + 10 + z.BlockHash.Msgsize() + 7
	if z.Signee == nil {
		s += msgp.NilSize
	} else {
		s += z.Signee.Msgsize()
	}
	s += 10
	if z.Signature == nil {
		s += msgp.NilSize
	} else {
		s += z.Signature.Msgsize()
	}
	return
}