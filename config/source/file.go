package source

import (
	"io/ioutil"
	"os"
)

// -
var (
	DefaultPath = "config.json"
)

// -
const (
	FILE = "file"
)

type file struct {
	path string
}

func (f *file) Read() (*ChangeSet, error) {
	fh, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	b, err := ioutil.ReadAll(fh)
	if err != nil {
		return nil, err
	}
	info, err := fh.Stat()
	if err != nil {
		return nil, err
	}

	cs := &ChangeSet{
		Timestamp: info.ModTime(),
		Checksum:  Sum(b),
		Data:      b,
	}

	return cs, nil
}

func (f *file) String() string {
	return FILE
}

// NewFile -
func NewFile(path string) Source {
	if path == "" {
		path = DefaultPath
	}

	return &file{path: path}
}
