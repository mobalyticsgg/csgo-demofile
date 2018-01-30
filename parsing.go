package demofile

import "github.com/Nyarum/barrel"

type Parser struct {
	header    *Header
	processor *barrel.Processor
}

func NewParser() *Parser {
	return &Parser{
		header: &Header{},
	}
}

func (p Parser) Parse(buf []byte) error {
	err := p.processor.WriteToBuffer(buf)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) parseHeader() error {

}
