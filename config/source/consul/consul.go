package consul

import (
	"bytes"
	"hawk/config/source"

	"github.com/axolotlteam/thunder/st"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

var (
	Clinet *api.Client
)

type consul struct {
	options    source.Options
	key        string
	configType string
	client     *api.Client
}

// Read -
func (c *consul) Read() (*source.ChangeSet, error) {
	v := viper.New()
	kv := c.client.KV()

	pair, _, err := kv.Get(c.key, nil)

	if err != nil || pair == nil {
		return nil, st.ErrorDataNotFound
	}
	switch c.configType {
	case "yaml":
		v.SetConfigType("yaml")
	case "json":
		v.SetConfigType("json")
	default:
		return nil, st.ErrorInvalidParameter
	}

	v.ReadConfig(bytes.NewReader(pair.Value))
	d, err := c.options.Encoder.Encode(v)
	cs := &source.ChangeSet{
		Data: d,
	}
	return cs, nil
}

// NewSource
func NewSource(opts ...source.Option) *consul {
	options := source.NewOptions(opts...)
	client := options.Context.Value(client{}).(*api.Client)
	key := options.Context.Value(key{}).(string)

	configtype := options.Context.Value(configType{}).(string)
	return &consul{
		client:     client,
		key:        key,
		configType: configtype,
	}
}

// GetClient -
func GetClient() *api.Client {
	return Clinet
}
