package demofile

import (
	"fmt"
	"io"
	"os"
)

// Demofile represents main info about demofile
type Demofile struct {
	file *os.File
	size int64
	buf  []byte

	parser *Parser

	isDebug bool
}

// NewDemofile creates Demofile structs with basis things
func NewDemofile(file string, isDebug bool) (*Demofile, error) {
	readFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	stat, err := readFile.Stat()
	if err != nil {
		return nil, err
	}

	return &Demofile{
		file:    readFile,
		size:    stat.Size(),
		buf:     make([]byte, stat.Size()),
		parser:  NewParser(isDebug),
		isDebug: isDebug,
	}, nil
}

// Start starts moving via demofile and execute events
func (d *Demofile) Start() error {
	defer d.file.Close()

	for {
		n, err := d.file.Read(d.buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		if d.isDebug {
			fmt.Println("Written bytes:", n)
		}

		bufToParse := make([]byte, n)
		copy(bufToParse, d.buf)

		err = d.parser.Parse(bufToParse)
		if err != nil {
			return err
		}
	}

	return nil
}
