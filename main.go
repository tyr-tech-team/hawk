package main

import (
	"hawk/registry/etcd"

	"github.com/davecgh/go-spew/spew"

	etcdnaming "github.com/coreos/etcd/clientv3/naming"

	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
)

func main() {
	ch := make(chan bool)

	etcdRegistry := etcd.NewRegistry(
		etcd.SetAddrs("127.0.0.1:2379"),
	)

	etcdRegistry.Registry()

	cli, _ := clientv3.NewFromURL("http://localhost:2379")

	r := &etcdnaming.GRPCResolver{Client: cli}

	b := grpc.RoundRobin(r)

	_, gerr := grpc.Dial("my-service", grpc.WithBalancer(b), grpc.WithInsecure())
	spew.Dump("!@@@@@")
	spew.Dump(gerr)

	<-ch
}
