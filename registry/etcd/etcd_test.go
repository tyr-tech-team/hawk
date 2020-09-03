package etcd

import "testing"

func TestReadEtcdRegistry(t *testing.T) {
	NewRegistry(
		SetAddrs("127.0.0.1:2379"),
	)

}
