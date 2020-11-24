package consul

import (
	"fmt"

	"github.com/tyr-tech-team/hawk/config/source"

	"github.com/hashicorp/consul/api"
)

// -
var (
	c *consul
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
		return nil, fmt.Errorf("not get any key/value")
	}

	cs := &source.ChangeSet{
		Data: pair.Value,
	}
	return cs, nil
}

// NewSource -
func NewSource(opts ...source.Option) (*consul, error) {
	options := source.NewOptions(opts...)
	key := options.Context.Value(key{}).(string)
	address := options.Context.Value(address{}).(string)
	configtype := options.Context.Value(configType{}).(string)

	config := &api.Config{
		Address: address,
	}

	cl, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	c = &consul{
		address:    address,
		client:     cl,
		key:        key,
		configType: configtype,
	}

	return c, nil
}

// Client -
func Client() *api.Client {
	return c.client
}
