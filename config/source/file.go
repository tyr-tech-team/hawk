package source

import (
	"io/ioutil"
	"os"

	"github.com/tyr-tech-team/hawk/config"
)

// -
var (
	DefaultPath = "config.json"
)

type file struct {
	path string
}

// NewFile -
func NewFile(path string) config.Source {
	if path == "" {
		path = DefaultPath
	}

	return &file{path: path}
}

func (f *file) Read() (*config.ChangeSet, error) {
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

	cs := &config.ChangeSet{
		Timestamp: info.ModTime(),
		Checksum:  Sum(b),
		Data:      b,
	}

	return cs, nil
}
