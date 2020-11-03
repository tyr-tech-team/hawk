package consul

import (
	"hawk/config/source"
	"log"

	"github.com/hashicorp/consul/api"
)

var (
	Clinet *api.Client
)

// kvRepository -
type consul struct {
	client  *api.Client
	options source.Options
}

// Read -
// func (c *consul) Read() (*source.ChangeSet, error) {

// }

// configure -
func configure(k *consul, opts *source.Options) {
	// connect consul agent
	// if len(opts.Context.Value(addressKey{})) != 1 {
	// 	log.Fatal("Addrs set failed")
	// }

	c := &api.Config{
		Address: opts.Context.Value(addressKey{}).(string),
	}

	client, err := api.NewClient(c)
	if err != nil {
		log.Fatal(err)
	}

	k.client = client

	Clinet = client
}

func NewSource(opts ...source.Option) *consul {
	options := source.NewOptions(opts...)

	k := &consul{
		options: options,
	}

	configure(k, &options)

	return k
}

// GetClient -
func GetClient() *api.Client {
	return Clinet
}

// String -
func (c *consul) String() string {
	return "consul"
}
