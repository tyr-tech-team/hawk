package consul

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hawk/config/source"

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
	address    string
	client     *api.Client
}

// Read -
func (c *consul) Read() (*source.ChangeSet, error) {
	v := viper.New()
	kv := c.client.KV()

	pair, _, err := kv.Get(c.key, nil)

	if err != nil || pair == nil {
		fmt.Println("kv error")
		return nil, fmt.Errorf("not get any key/value")
	}
	fmt.Println("kv", kv)
	switch c.configType {
	case "yaml":
		v.SetConfigType("yaml")
	case "json":
		v.SetConfigType("json")
	default:
		return nil, fmt.Errorf("Can't found configType")
	}

	v.ReadConfig(bytes.NewReader(pair.Value))

	// d, err := c.options.Encoder.Encode(v.Unmarshal())
	fmt.Print("v ", v)
	ans, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
	}

	cs := &source.ChangeSet{
		Data: ans,
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
