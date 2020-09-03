package file

import (
	"hawk/config/source"
	"io/ioutil"
	"os"
)

var (
	DefaultPath = "config.json"
)

type file struct {
	path string
	opts source.Options
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
		Format:    format(f.path, f.opts.Encoder),
		Source:    f.String(),
		Timestamp: info.ModTime(),
		Data:      b,
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

func (f *file) String() string {
	return "file"
}

func NewSource(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)
	path := DefaultPath

	f, ok := options.Context.Value(filePathKey{}).(string)
	if ok {
		path = f
	}

	return &file{opts: options, path: path}
}

