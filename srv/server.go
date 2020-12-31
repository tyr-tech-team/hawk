package srv

import (
	"net"
)

// Server -
type Server interface {
}

type server struct {
	register Register
	host     string
	listener net.Listener
}

// New -
func New() Server {

	//listener, err := net.Listen("tcp", fmt.Sprintf(""))

	return nil
}

func (s *server) Listener() net.Listener {
	return s.listener
}
