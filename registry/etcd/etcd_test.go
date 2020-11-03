package etcd

import (
	"testing"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
)

func TestReadEtcdRegistry(t *testing.T) {
	etcdRegistry := NewRegistry(
		SetAddrs("127.0.0.1:2379"),
	)
	etcdRegistry.Registry()

	cli, _ := clientv3.NewFromURL("http://localhost:2379")
	r := &etcdnaming.GRPCResolver{Client: cli}

	b := grpc.RoundRobin(r)

	conn, gerr := grpc.Dial("my-service", grpc.WithBalancer(b))
	spew.Dump("!@@@@@")
	spew.Dump(conn)
	spew.Dump(gerr)

}

