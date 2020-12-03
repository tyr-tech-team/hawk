package file

import (
	"io/ioutil"
	"os"

	"github.com/tyr-tech-team/hawk/config/source"
	"github.com/tyr-tech-team/hawk/encoder"
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
	path    string
	encoder encoder.Encoder
}

func (f *file) Read() (*source.ChangeSet, error) {
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

	cs := &source.ChangeSet{
		Format:    format(f.path, f.encoder),
		Source:    f.String(),
		Timestamp: info.ModTime(),
		Data:      b,
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

func (f *file) String() string {
	return FILE
}

// NewSource -
func NewSource(path string, e encoder.Encoder) source.Source {
	if path == "" {
		path = DefaultPath
	}

	return &file{path: path}
}
