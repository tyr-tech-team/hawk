package consul

import (
	"log"

	"github.com/hashicorp/consul/api"
)

// Config -
type Config struct {
	Address   string `json:"url"`
	ACL       string `json:"acl"`
	Namespace string `json:"namespace"`
}

// Client -
type Client interface {
}

type client struct {
	consul *api.Client
}

// NewClient -
func NewClient(c Config) Client {
	cc := api.Config{
		Address: c.Address,
		Token:   c.ACL,
	}

	conn, err := api.NewClient(&cc)
	if err != nil {
		log.Panic(err)
	}

	return &client{
		consul: conn,
	}

}
