package demofile

import (
	"io"
	"os"
)

type Demofile struct {
	file io.Reader
}

func NewDemofile(file string) (*Demofile, error) {
	readFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return &Demofile{
		file: readFile,
	}, nil
}

func (d *Demofile) Start() error {
	return nil
}
