package demofile

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Nyarum/barrel"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/geo/r3"
)

const (
	minSizeToParse = 2
)

var (
	errNotSufficientBytes = errors.New("Not sufficient bytes for parse chunk")
)

type Parser struct {
	header  *Header
	packet  *Packet
	cmdInfo *CmdInfo
	chunk   *Chunk

	processor *barrel.Processor

	isDebug bool
}

func NewParser(isDebug bool) *Parser {
	return &Parser{
		processor: barrel.NewProcessor([]byte{}).SetEndian(barrel.LittleEndian),
		packet:    &Packet{},
		cmdInfo:   &CmdInfo{},
		chunk:     &Chunk{},
		isDebug:   isDebug,
	}
}

func (p *Parser) Parse(buf []byte) error {
	err := p.processor.WriteToBuffer(buf)
	if err != nil {
		return err
	}

	if p.header == nil {
		err = p.parseHeader()
		if err != nil {
			return err
		}

		if p.isDebug {
			spew.Dump(p.header)
		}
	}

	for {
		if p.processor.Buffer().Len() < minSizeToParse {
			break
		}

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
	p.processor.
		ReadInt8(&p.packet.Cmd).
		ReadInt32(&p.packet.Tick).
		ReadInt8(&p.packet.PlayerSlot)

	if p.processor.Error() != nil {
		return p.processor.Error()
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
	var (
		x, y, z float32
	)

	p.processor.
		ReadFloat32(&x).
		ReadFloat32(&y).
		ReadFloat32(&z)

	v.X = float64(x)
	v.Y = float64(y)
	v.Z = float64(z)

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
	err = p.parseChunk(cmdPacket)
	if err != nil {
		return err
	}

	return p.processor.Error()
}

func (p *Parser) parseChunk(cmd command) error {
	p.processor.ReadInt32(&p.chunk.Lenght)

	if p.processor.Error() != nil {
		return p.processor.Error()
	}

	if p.processor.Buffer().Len() < int(p.chunk.Lenght) {
		return errNotSufficientBytes
	}

	// Skip commands
	if cmd == cmdUser || cmd == cmdConsole || cmd == cmdPacket {
		p.processor.Skip(int(p.chunk.Lenght))

		return p.processor.Error()
	}

	p.processor.ReadBytes(&p.chunk.Data, int(p.chunk.Lenght))

	if cmd == cmdStringTables {
		fmt.Println(cmd)
		fmt.Println(p.packet.Tick)

		err := p.parseStringTables()
		if err != nil {
			return err
		}
	}

	return p.processor.Error()
}

func (p *Parser) parseStringTables() error {
	pr := barrel.NewProcessor(p.chunk.Data)

	var numTables int8
	pr.ReadInt8(&numTables)
	if pr.Error() != nil {
		return pr.Error()
	}

	for i := 0; i < int(numTables); i++ {
		var tableName string
		pr.ReadStringEOF(&tableName)

		if p.isDebug {
			fmt.Println("Table name:", tableName)
		}

		err := p.parseStringTable(pr, tableName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseStringTable(pr *barrel.Processor, tableName string) error {
	var numStrings int16
	pr.ReadInt16(&numStrings)

	for i := 0; i < int(numStrings); i++ {
		var (
			stringName string
			pass       bool
		)

		pr.ReadStringEOF(&stringName)
		pr.ReadBit(&pass)

		fmt.Println(stringName)

		if pass {
			var dataSize uint16
			pr.ReadUint16(&dataSize)

			data := make([]byte, int(dataSize))
			pr.ReadBytes(&data, int(dataSize))
		}
	}

	var clientStaffPass bool
	pr.ReadBit(&clientStaffPass)

	if clientStaffPass {
		var numStrings uint16
		pr.ReadUint16(&numStrings)

		for i := 0; i < int(numStrings); i++ {
			var (
				stringName string
				pass       bool
			)

			pr.ReadStringEOF(&stringName)
			pr.ReadBit(&pass)

			if pass {
				var numFields uint16
				pr.ReadUint16(&numFields)
				pr.Skip(int(numFields))
			}

		}
	}

	return pr.Error()
}
