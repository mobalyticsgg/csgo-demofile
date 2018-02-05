package bitparser

import (
	"bytes"
	"math"

	bitread "github.com/markus-wa/gobitread"
)

const (
	numBitsInByte = 8
)

type Bitparser struct {
	*bitread.BitReader
	tempBuf *bytes.Buffer
}

func NewBitparser(buf []byte) *Bitparser {
	br := &bitread.BitReader{}
	br.Open(bytes.NewBuffer(buf), 4096)

	return &Bitparser{
		BitReader: br,
		tempBuf:   bytes.NewBuffer(make([]byte, 4096)),
	}
}

func (b *Bitparser) ReadStringWithLen(n int) (string, error) {
	return b.ReadCString(n), nil
}

func (b *Bitparser) ReadStringEOF() (string, error) {
	for {
		v := b.ReadSingleByte()
		if v == 0 {
			break
		}

		b.tempBuf.WriteByte(v)
	}

	defer b.tempBuf.Reset()
	return b.tempBuf.String(), nil
}

func (b *Bitparser) ReadFloat32() (float32, error) {
	v, err := b.ReadUint32()
	if err != nil {
		return 0, err
	}

	return math.Float32frombits(v), nil
}

func (b *Bitparser) ReadUint16() (uint16, error) {
	v, err := b.ReadUint16()
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (b *Bitparser) ReadUint32() (uint32, error) {
	v, err := b.ReadUint32()
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (b *Bitparser) ReadInt32() (int32, error) {
	v, err := b.ReadInt32()
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (b *Bitparser) Skip(n int) error {
	b.ReadBytes(n)

	return nil
}

// FIX
/*
// Bit reads the next bit as a boolean.
func (r *Reader) Bit() bool {
	res := (r.src[r.idx>>3] & (1 << uint(r.idx&7)))
	r.idx++
	return res != 0
}
*/
