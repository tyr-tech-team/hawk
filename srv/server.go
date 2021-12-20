package srv

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/tyr-tech-team/hawk/config"
	"github.com/tyr-tech-team/hawk/tools/network"
)

// Server -
type Server interface {
	GetListener() net.Listener
	Register() error
	Deregister() error
	GetHost() string
	Close()
}

type server struct {
	name                  string
	host                  string
	port                  string
	protocol              config.Protocol
	register              Register
	traefik               bool
	serviceRegisterConfig config.ServiceRegister
	listener              net.Listener
}

// New -
func New(opts ...Options) Server {
	s := &server{}
	for _, v := range opts {
		v(s)
	}

	s.newTCPListener()
	s.setServiceRegister()

	return *s
}

// Listener -
func (s server) GetListener() net.Listener {
	return s.listener
}

func (s server) GetHost() string {
	return fmt.Sprintf("%s:%s", s.host, s.port)
}

// Close -
func (s server) Close() {
	s.listener.Close()
}

// Register -  註冊服務
func (s server) Register() error {
	s.register.SetRegisterConfig(s.serviceRegisterConfig)
	return s.register.Register()
}

// Deregister - 取消註冊服務
func (s server) Deregister() error {
	return s.register.Deregister()
}

func (s *server) SetName(name string) {
	s.name = name
}

func (s *server) SetHost(host string) {
	s.host = host
	if host == "" {
		s.host = network.LocalIP()
	}
}

func (s *server) SetRegister(r Register) {
	s.register = r
}

func (s *server) SetPort(port string) {
	s.port = port
}

func (s *server) SetProtocal(p config.Protocol) {
	s.protocol = p
}

func (s *server) SetTraefik(setup bool) {
	s.traefik = setup
}

func (s *server) setServiceRegister() {
	port, err := strconv.Atoi(s.port)
	if err != nil {
		panic(err)
	}

	s.serviceRegisterConfig = config.ServiceRegister{
		Name:    s.name,
		Address: s.host,
		Port:    port,
		Traefik: s.traefik,
		Protocol: func() config.Protocol {
			if s.protocol == "" {
				return config.HTTP
			}
			return s.protocol
		}(),
	}
}

func (s *server) newTCPListener() {
	list, err := net.Listen("tcp4", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		panic(err)
	}

	s.listener = list

	host := strings.Split(list.Addr().String(), ":")
	s.host = host[0]
	s.port = host[len(host)-1]
}
