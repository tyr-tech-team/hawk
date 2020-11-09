package consul

import (
	"fmt"
	"hawk/config/source"

	"github.com/hashicorp/consul/api"
)

var (
	Clinet *api.Client
)

type consul struct {
	options    source.Options
	key        string
	configType string
	address    string
	client     *api.Client
}

// Read -
func (c *consul) Read() (*source.ChangeSet, error) {
	kv := c.client.KV()

	pair, _, err := kv.Get(c.key, nil)

	if err != nil || pair == nil {
		fmt.Println("kv error")
		return nil, fmt.Errorf("not get any key/value")
	}

	cs := &source.ChangeSet{
		Data: pair.Value,
	}
	return cs, nil
}

// NewSource
func NewSource(opts ...source.Option) *consul {
	options := source.NewOptions(opts...)
	key := options.Context.Value(key{}).(string)
	address := options.Context.Value(address{}).(string)
	configtype := options.Context.Value(configType{}).(string)

	c := &api.Config{
		Address: address,
	}

	client, err := api.NewClient(c)

	if err != nil {
		fmt.Errorf(err.Error())
	}
	return &consul{
		address:    address,
		client:     client,
		key:        key,
		configType: configtype,
	}
}

// GetClient -
func GetClient() *api.Client {
	return Clinet
}
