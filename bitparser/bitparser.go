package bitparser

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/bamiaux/iobit"
)

type Bitparser struct {
	iobit.Reader
	tempBuf *bytes.Buffer
}

func NewBitparser(buf []byte) *Bitparser {
	return &Bitparser{
		Reader:  iobit.NewReader(buf),
		tempBuf: bytes.NewBuffer(make([]byte, 4096)),
	}
}

func (b *Bitparser) ReadStringEOF() string {
	for {
		var v byte
		if v = b.Byte(); v == 0 || v == 10 {
			break
		}

		b.tempBuf.WriteByte(v)
	}

	defer b.tempBuf.Reset()
	return b.tempBuf.String()
}

func (b *Bitparser) Float32() float32 {
	return math.Float32frombits(b.Le32())
}

func (b *Bitparser) LInt32() int32 {
	buf := b.Bytes(4)

	return int32(binary.LittleEndian.Uint32(buf))
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
