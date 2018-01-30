package demofile

import (
	"fmt"
	"io"
	"os"
)

const (
	maxSizeDelim = 10
)

type Demofile struct {
	file io.Reader
	size int64
	buf  []byte

	parser *Parser
}

func NewDemofile(file string) (*Demofile, error) {
	readFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	stat, err := readFile.Stat()
	if err != nil {
		return nil, err
	}

	return &Demofile{
		file:   readFile,
		size:   stat.Size(),
		buf:    make([]byte, stat.Size()/maxSizeDelim),
		parser: NewParser(),
	}, nil
}

func (d *Demofile) Start() error {
	for {
		n, err := d.file.Read(d.buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		success, err := d.parser.Parse(d.buf)
		if err != nil {
			return err
		}

		if !success {
			fmt.Println("Whoops")
		}

		fmt.Println("Written", n)
	}

	return nil
}
