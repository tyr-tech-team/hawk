package file

import (
	"context"
	"hawk/config/source"

	"github.com/davecgh/go-spew/spew"
)

type filePathKey struct{}

// WithPath sets the path to file
func WithPath(p string) source.Option {
	spew.Dump("WithPath")
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, filePathKey{}, p)
	}
}

