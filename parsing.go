package demofile

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MobalyticsGG/csgo-demofile/bitparser"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/geo/r3"
)

const (
	minSizeToParse = 2
	numBitsInByte  = 8
)

var (
	errNotSufficientBytes = errors.New("Not sufficient bytes for parse chunk")
)

type Parser struct {
	header  *Header
	packet  *Packet
	cmdInfo *CmdInfo
	chunk   *Chunk

	bitparser *bitparser.Bitparser

	isDebug bool
}

func NewParser(isDebug bool) *Parser {
	return &Parser{
		packet:  &Packet{},
		cmdInfo: &CmdInfo{},
		chunk:   &Chunk{},
		isDebug: isDebug,
	}
}

func (p *Parser) Parse(buf []byte) error {
	if p.bitparser == nil {
		p.bitparser = bitparser.NewBitparser(buf)
	}

	if p.header == nil {
		err := p.parseHeader()
		if err != nil {
			return err
		}

		if p.isDebug {
			spew.Dump(p.header)
		}
	}

	for {
		err := p.handlePacket()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) handlePacket() error {
	var err error

	p.packet.Cmd = p.bitparser.ReadSingleByte()
	p.packet.Tick, err = p.bitparser.ReadInt32()
	p.packet.PlayerSlot = p.bitparser.ReadSingleByte()

	if err != nil {
		return err
	}

	cmd := command(p.packet.Cmd)
	switch cmd {
	case cmdSignOn, cmdPacket:
		err := p.parseCmdInfo()
		if err != nil {
			return err
		}

	case cmdSynctick:
		break

	case cmdConsole, cmdUser, cmdDataTables, cmdStringTables:
		err := p.parseChunk(cmd)
		if err != nil {
			return err
		}

	case cmdStop:
		return nil

	case cmdCustomData:
		if p.isDebug {
			fmt.Println("Found custom data but we don't have any logic for that")
		}
	}

	return nil
}

func (p *Parser) parseHeader() error {
	p.header = &Header{}

	var err error
	p.header.Filestamp, err = p.bitparser.ReadStringWithLen(maxFilestampSize)
	p.header.Protocol, err = p.bitparser.ReadUint32()
	p.header.NetworkProtocol, err = p.bitparser.ReadUint32()
	p.header.ServerName, err = p.bitparser.ReadStringWithLen(maxOsPath)
	p.header.ClientName, err = p.bitparser.ReadStringWithLen(maxOsPath)
	p.header.MapName, err = p.bitparser.ReadStringWithLen(maxOsPath)
	p.header.GameDirectory, err = p.bitparser.ReadStringWithLen(maxOsPath)
	p.header.PlaybackTime, err = p.bitparser.ReadFloat32()
	p.header.PlaybackTicks, err = p.bitparser.ReadUint32()
	p.header.PlaybackFrames, err = p.bitparser.ReadUint32()
	p.header.SignOnLenght, err = p.bitparser.ReadUint32()

	if err != nil {
		return err
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
	var (
		x, y, z float32
		err     error
	)
	x, err = p.bitparser.ReadFloat32()
	y, err = p.bitparser.ReadFloat32()
	z, err = p.bitparser.ReadFloat32()

	if err != nil {
		return err
	}

	v.X, v.Y, v.Z = float64(x), float64(y), float64(z)

	return nil
}

func (p *Parser) parseCmdInfo() error {
	var err error

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
		splitCmdInfo.Flags, err = p.bitparser.ReadInt32()
		err := originViewParse(&splitCmdInfo.Original)
		if err != nil {
			return err
		}

		return originViewParse(&splitCmdInfo.Resampled)
	}

	err = partCmdParse(&p.cmdInfo.Parts[0])
	if err != nil {
		return err
	}

	err = partCmdParse(&p.cmdInfo.Parts[1])
	if err != nil {
		return err
	}

	p.bitparser.Skip(8)
	err = p.parseChunk(cmdPacket)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) parseChunk(cmd command) error {
	var err error

	p.chunk.Lenght, err = p.bitparser.ReadUint32()
	if err != nil {
		return err
	}

	// Skip commands
	if cmd == cmdUser || cmd == cmdConsole || cmd == cmdPacket {
		return p.bitparser.Skip(int(p.chunk.Lenght))
	}

	p.chunk.Data = p.bitparser.ReadBytes(int(p.chunk.Lenght))

	if cmd == cmdStringTables {
		fmt.Println(cmd)
		fmt.Println(p.packet.Tick)

		err := p.parseStringTables()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseStringTables() error {
	br := bitparser.NewBitparser(p.chunk.Data)

	numTables := br.ReadSingleByte()
	for i := 0; i < int(numTables); i++ {
		tableName, err := br.ReadStringEOF()
		if err != nil {
			return err
		}

		if p.isDebug {
			fmt.Println("Table name:", tableName)
		}

		err = p.parseStringTable(br, tableName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseStringTable(br *bitparser.Bitparser, tableName string) error {
	numStrings, err := br.ReadUint16()
	if err != nil {
		return err
	}

	fmt.Println("num:", numStrings)
	for i := 0; i < int(numStrings); i++ {
		stringName, err := br.ReadStringEOF()
		if err != nil {
			return err
		}
		fmt.Println(stringName)

		pass := br.ReadBit()

		fmt.Println(pass)
		if pass {
			dataSize, err := br.ReadUint16()
			if err != nil {
				return err
			}

			fmt.Println(dataSize)
			br.ReadBytes(int(dataSize))
		}
	}

	pass := br.ReadBit()

	if pass {
		numStrings, err := br.ReadUint16()
		if err != nil {
			return err
		}

		for i := 0; i < int(numStrings); i++ {
			br.ReadStringEOF()

			pass := br.ReadBit()

			if pass {
				numFields, err := br.ReadUint16()
				if err != nil {
					return err
				}

				br.Skip(int(numFields))
			}

		}
	}

	return nil
}
