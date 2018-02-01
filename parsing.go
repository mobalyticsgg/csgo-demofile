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
			if err == errNotSufficientBytes {
				break
			}

			return err
		}
	}

	return nil
}

func (p *Parser) handlePacket() error {
	p.packet.Cmd = p.bitparser.Byte()
	p.packet.Tick = p.bitparser.LInt32()
	p.packet.PlayerSlot = p.bitparser.Byte()

	if p.bitparser.Error() != nil {
		return p.bitparser.Error()
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

	p.header.Filestamp = p.bitparser.String(maxFilestampSize)
	p.header.Protocol = p.bitparser.Le32()
	p.header.NetworkProtocol = p.bitparser.Le32()
	p.header.ServerName = p.bitparser.String(maxOsPath)
	p.header.ClientName = p.bitparser.String(maxOsPath)
	p.header.MapName = p.bitparser.String(maxOsPath)
	p.header.GameDirectory = p.bitparser.String(maxOsPath)
	p.header.PlaybackTime = p.bitparser.Float32()
	p.header.PlaybackTicks = p.bitparser.Le32()
	p.header.PlaybackFrames = p.bitparser.Le32()
	p.header.SignOnLenght = p.bitparser.Le32()

	if p.bitparser.Error() != nil {
		return p.bitparser.Error()
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
	v.X = float64(p.bitparser.Float32())
	v.Y = float64(p.bitparser.Float32())
	v.Z = float64(p.bitparser.Float32())

	return p.bitparser.Error()
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
		splitCmdInfo.Flags = p.bitparser.Int32(32)
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

	p.bitparser.Skip(8 * numBitsInByte)
	err = p.parseChunk(cmdPacket)
	if err != nil {
		return err
	}

	return p.bitparser.Error()
}

func (p *Parser) parseChunk(cmd command) error {
	p.chunk.Lenght = p.bitparser.Le32()
	if p.bitparser.Error() != nil {
		return p.bitparser.Error()
	}

	// Skip commands
	if cmd == cmdUser || cmd == cmdConsole || cmd == cmdPacket {
		p.bitparser.Skip(uint(p.chunk.Lenght * numBitsInByte))

		return p.bitparser.Error()
	}

	p.chunk.Data = p.bitparser.Bytes(int(p.chunk.Lenght))

	if cmd == cmdStringTables {
		fmt.Println(cmd)
		fmt.Println(p.packet.Tick)

		err := p.parseStringTables()
		if err != nil {
			return err
		}
	}

	return p.bitparser.Error()
}

func (p *Parser) parseStringTables() error {
	br := bitparser.NewBitparser(p.chunk.Data)

	numTables := br.Byte()
	if br.Error() != nil {
		return br.Error()
	}

	for i := 0; i < int(numTables); i++ {
		tableName := br.ReadStringEOF()
		if p.isDebug {
			fmt.Println("Table name:", tableName)
		}

		err := p.parseStringTable(br, tableName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseStringTable(br *bitparser.Bitparser, tableName string) error {
	numStrings := br.Le16()
	fmt.Println("num:", numStrings)
	for i := 0; i < int(numStrings); i++ {
		stringName := br.ReadStringEOF()
		fmt.Println(stringName)

		pass := br.Bit()
		fmt.Println(pass)
		if pass {
			dataSize := br.Le16()
			fmt.Println(dataSize)
			br.Bytes(int(dataSize))
		}
	}

	if br.Bit() {
		numStrings := br.Le16()
		for i := 0; i < int(numStrings); i++ {
			br.ReadStringEOF()
			if br.Bit() {
				numFields := br.Le16()
				br.Skip(uint(numFields * numBitsInByte))
			}

		}
	}

	return br.Error()
}
