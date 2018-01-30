package demofile

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Nyarum/barrel"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/geo/r3"
)

type Parser struct {
	header  *Header
	packet  *Packet
	cmdInfo *CmdInfo
	chunk   *Chunk

	processor *barrel.Processor
}

func NewParser() *Parser {
	return &Parser{
		processor: barrel.NewProcessor([]byte{}).SetEndian(barrel.LittleEndian),
		packet:    &Packet{},
		cmdInfo:   &CmdInfo{},
		chunk:     &Chunk{},
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

func (p *Parser) handlePacket() (bool, error) {
	p.processor.
		ReadInt8(&p.packet.Cmd).
		ReadInt32(&p.packet.Tick).
		ReadInt8(&p.packet.PlayerSlot)

	if p.processor.Error() != nil {
		return false, p.processor.Error()
	}

	cmd := command(p.packet.Cmd)
	switch cmd {
	case cmdSignOn, cmdPacket:
		err := p.parseCmdInfo()
		if err != nil {
			return false, p.processor.Error()
		}

	case cmdSynctick:
		return false, errors.New("")

	case cmdConsole, cmdUser, cmdDataTables, cmdStringTables:
		err := p.parseChunk(cmd)
		if err != nil {
			return false, p.processor.Error()
		}

	case cmdStop:
		return false, nil

	case cmdCustomData:
		fmt.Println("Found custom data but we don't have any logic for that")
	}

	return true, nil
}

func (p *Parser) parseHeader() error {
	p.header = &Header{}

	p.processor.
		ReadStringWithLen(maxFilestampSize, &p.header.Filestamp).
		ReadInt32(&p.header.Protocol).
		ReadInt32(&p.header.NetworkProtocol).
		ReadStringWithLen(maxOsPath, &p.header.ServerName).
		ReadStringWithLen(maxOsPath, &p.header.ClientName).
		ReadStringWithLen(maxOsPath, &p.header.MapName).
		ReadStringWithLen(maxOsPath, &p.header.GameDirectory).
		ReadFloat32(&p.header.PlaybackTime).
		ReadInt32(&p.header.PlaybackTicks).
		ReadInt32(&p.header.PlaybackFrames).
		ReadInt32(&p.header.SignOnLenght)

	if p.processor.Error() != nil {
		return p.processor.Error()
	}

	// Remove \x00 from strings
	p.header.Filestamp = strings.Trim(p.header.Filestamp, "\x00")
	p.header.ServerName = strings.Trim(p.header.ServerName, "\x00")
	p.header.ClientName = strings.Trim(p.header.ClientName, "\x00")
	p.header.MapName = strings.Trim(p.header.MapName, "\x00")
	p.header.GameDirectory = strings.Trim(p.header.GameDirectory, "\x00")

	return nil
}

func (p *Parser) parseVector(v *r3.Vector) error {
	p.processor.
		ReadFloat64(&v.X).
		ReadFloat64(&v.Y).
		ReadFloat64(&v.Z)

	return p.processor.Error()
}

func (p *Parser) parseCmdInfo() error {
	originViewParse := func(originView *OriginViewAngles) error {
		err := p.parseVector(&originView.ViewOrigin)
		if err != nil {
			return err
		}

		err = p.parseVector(&originView.ViewAngles)
		if err != nil {
			return err
		}

		return p.parseVector(&originView.LocalViewAngles)
	}

	partCmdParse := func(splitCmdInfo *SplitCmdInfo) error {
		p.processor.ReadInt32(&splitCmdInfo.Flags)

		err := originViewParse(&splitCmdInfo.Original)
		if err != nil {
			return err
		}

		return originViewParse(&splitCmdInfo.Resampled)
	}

	err := partCmdParse(&p.cmdInfo.Parts[0])
	if err != nil {
		return err
	}

	err = partCmdParse(&p.cmdInfo.Parts[1])
	if err != nil {
		return err
	}

	p.processor.Skip(8)
	p.parseChunk(cmdPacket)

	return p.processor.Error()
}

func (p *Parser) parseChunk(cmd command) error {
	p.processor.
		ReadInt32(&p.chunk.Lenght)

	if p.processor.Error() != nil {
		return p.processor.Error()
	}

	// Skip commands
	if cmd == cmdUser || cmd == cmdConsole || cmd == cmdPacket {
		p.processor.Skip(int(p.chunk.Lenght))

		return p.processor.Error()
	}

	p.processor.ReadBytes(&p.chunk.Data, int(p.chunk.Lenght))

	return p.processor.Error()
}
