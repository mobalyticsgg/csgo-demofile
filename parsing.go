package demofile

import (
	"strings"

	"github.com/Nyarum/barrel"
	"github.com/davecgh/go-spew/spew"
)

type Parser struct {
	header    *Header
	processor *barrel.Processor
}

func NewParser() *Parser {
	return &Parser{
		processor: barrel.NewProcessor([]byte{}).SetEndian(barrel.LittleEndian),
	}
}

func (p *Parser) Parse(buf []byte) (bool, error) {
	err := p.processor.WriteToBuffer(buf)
	if err != nil {
		return false, err
	}

	if p.header == nil {
		err = p.parseHeader()
		if err != nil {
			return false, err
		}

		spew.Dump(p.header)
	}

	return true, nil
}

func (p *Parser) parseHeader() error {
	p.header = &Header{}

	p.processor.ReadStringWithLen(maxFilestampSize, &p.header.Filestamp)
	p.processor.ReadInt32(&p.header.Protocol)
	p.processor.ReadInt32(&p.header.NetworkProtocol)
	p.processor.ReadStringWithLen(maxOsPath, &p.header.ServerName)
	p.processor.ReadStringWithLen(maxOsPath, &p.header.ClientName)
	p.processor.ReadStringWithLen(maxOsPath, &p.header.MapName)
	p.processor.ReadStringWithLen(maxOsPath, &p.header.GameDirectory)
	p.processor.ReadFloat32(&p.header.PlaybackTime)
	p.processor.ReadInt32(&p.header.PlaybackTicks)
	p.processor.ReadInt32(&p.header.PlaybackFrames)
	p.processor.ReadInt32(&p.header.SignOnLenght)

	// Remove \x00 from strings
	p.header.Filestamp = strings.Trim(p.header.Filestamp, "\x00")
	p.header.ServerName = strings.Trim(p.header.ServerName, "\x00")
	p.header.ClientName = strings.Trim(p.header.ClientName, "\x00")
	p.header.MapName = strings.Trim(p.header.MapName, "\x00")
	p.header.GameDirectory = strings.Trim(p.header.GameDirectory, "\x00")

	return p.processor.Error()
}

func (p *Parser) parseCmdInfo() error {
	return nil
}
