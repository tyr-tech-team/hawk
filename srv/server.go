package srv

import "net"

// Server -
type Server interface {
}

type server struct {
	register Register
	listener net.Listener
}

// New -
func New() Server {
	return nil
}
