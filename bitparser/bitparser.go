package bitparser

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/icza/bitio"
)

const (
	numBitsInByte = 8
)

type Bitparser struct {
	bitio.Reader
	tempBuf *bytes.Buffer
}

func NewBitparser(buf []byte) *Bitparser {
	return &Bitparser{
		Reader:  bitio.NewReader(bytes.NewBuffer(buf)),
		tempBuf: bytes.NewBuffer(make([]byte, 4096)),
	}
}

func (b *Bitparser) ReadStringWithLen(n int) (string, error) {
	buf := make([]byte, int(n))
	_, err := b.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (b *Bitparser) ReadStringEOF() (string, error) {
	for {
		v, err := b.ReadByte()
		if err != nil {
			return "", err
		}

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
	buf := make([]byte, 2)
	_, err := b.Read(buf)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint16(buf), nil
}

func (b *Bitparser) ReadUint32() (uint32, error) {
	buf := make([]byte, 4)
	_, err := b.Read(buf)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(buf), nil
}

func (b *Bitparser) ReadInt32() (int32, error) {
	buf := make([]byte, 4)
	_, err := b.Read(buf)
	if err != nil {
		return 0, err
	}

	return int32(binary.LittleEndian.Uint32(buf)), nil
}

func (b *Bitparser) Skip(n int) error {
	buf := make([]byte, n)
	_, err := b.Read(buf)

	return err
}

func (b *Bitparser) ReadBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := b.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
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
