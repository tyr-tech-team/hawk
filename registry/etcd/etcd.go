package etcd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"hawk/registry"
	"log"
	"net"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc/naming"
)

var (
	DefaultEndpoints = []string{"127.0.0.1:2379"}
	DefaultTimeout   = 30 * time.Second
	DefaultPrefix    = "/registry/"
)

type etcdRegistry struct {
	client  *clientv3.Client
	options registry.Options

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

	// set timeout
	if e.options.Timeout == 0 {
		e.options.Timeout = DefaultTimeout
	}

	// set username, password
	if e.options.Context != nil {
		u, ok := e.options.Context.Value(authKey{}).(*authCreds)
		if ok {
			config.Username = u.Username
			config.Password = u.Password
		}
	}

	if e.options.Secure || e.options.TLSConfig != nil {
		tlsConfig := e.options.TLSConfig
		if tlsConfig == nil {
			tlsConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}

		config.TLS = tlsConfig
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

	return nil
}

// Registry -
func (e *etcdRegistry) Registry() error {
	val := &naming.Update{
		Op:   naming.Add,
		Addr: "192.168.7.7",
	}

	grant, _ := e.client.Grant(context.TODO(), 5)

	e.client.Put(context.TODO(), DefaultPrefix, encode(val), clientv3.WithLease(grant.ID))

	go e.keepAlive(grant.ID)

	return nil
}

func (e *etcdRegistry) keepAlive(id clientv3.LeaseID) error {
	ch, err := e.client.KeepAlive(context.TODO(), id)
	if err != nil {
		log.Fatalf("etcd keepAlive failed : %v", err)
		return err
	}

	for {
		select {
		case resp, ok := <-ch:
			if !ok {
				e.Revoke()
				return nil
			}
			spew.Dump(resp, ok)
		}
	}
}

// Revoke -
func (e *etcdRegistry) Revoke() error {
	return nil
}

func encode(n *naming.Update) string {
	b, _ := json.Marshal(n)
	return string(b)
}

func decode(n []byte) *naming.Update {
	var s *naming.Update
	json.Unmarshal(n, &s)
	return s
}
