package etcd

import (
	"context"
	"hawk/registry"
	"net"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/davecgh/go-spew/spew"
)

var (
	DefaultEndpoints = []string{"127.0.0.1:2379"}
)

type etcdRegistry struct {
	client  *clientv3.Client
	options registry.Options

	sync.RWMutex
	register map[string]uint64
	leases   map[string]clientv3.LeaseID
}

// NewRegistry -
func NewRegistry(opts ...registry.Option) registry.Registry {
	e := &etcdRegistry{
		options:  registry.Options{},
		register: make(map[string]uint64),
		leases:   make(map[string]clientv3.LeaseID),
	}

	configure(e, opts...)

	return e
}

func configure(e *etcdRegistry, opts ...registry.Option) error {
	config := clientv3.Config{
		Endpoints: DefaultEndpoints,
	}

	for _, o := range opts {
		o(&e.options)
	}

	if e.options.Timeout == 0 {
		e.options.Timeout = 5 * time.Second
	}

	if e.options.Context != nil {
		u, ok := e.options.Context.Value(authKey{}).(*authCreds)
		if ok {
			config.Username = u.Username
			config.Password = u.Password
		}
	}

	var cAddrs []string

	for _, address := range e.options.Addrs {
		if len(address) == 0 {
			continue
		}
		addr, port, err := net.SplitHostPort(address)
		if ae, ok := err.(*net.AddrError); ok && ae.Err == "missing port in address" {
			port = "2379"
			addr = address
			cAddrs = append(cAddrs, net.JoinHostPort(addr, port))
		} else if err == nil {
			cAddrs = append(cAddrs, net.JoinHostPort(addr, port))
		}
	}

	// if we got addrs then we'll update
	if len(cAddrs) > 0 {
		config.Endpoints = cAddrs
	}

	cli, err := clientv3.New(config)
	if err != nil {
		return err
	}

	e.client = cli

	resp, _ := cli.Grant(context.TODO(), 5)
	spew.Dump(resp.ID)

	key := "/micro/registry/"
	cli.Put(context.TODO(), key, "192.168.7.7", clientv3.WithLease(resp.ID))
	ch, _ := cli.KeepAlive(context.TODO(), resp.ID)

	for {
		select {
		case resp, ok := <-ch:
			spew.Dump(resp, ok)
		}
	}

	return nil
}
