package srv

import "github.com/tyr-tech-team/hawk/config"

// Options -
type Options func(*server)

// SetName -
func SetName(name string) Options {
	return func(s *server) {
		s.SetName(name)
	}
}

// SetHost -
func SetHost(host string) Options {
	return func(s *server) {
		s.SetHost(host)
	}
}

// SetPort -
func SetPort(port string) Options {
	return func(s *server) {
		s.SetPort(port)

	}
}

// SetRegister -
func SetRegister(r Register) Options {
	return func(s *server) {
		s.SetRegister(r)
	}
}

// SetEnableTraefik -
func SetEnableTraefik() Options {
	return func(s *server) {
		s.SetTraefik(true)
	}
}

// SetGRPC -
func SetGRPC() Options {
	return func(s *server) {
		s.SetProtocal(config.GRPC)
	}
}

// SetHTTP -
func SetHTTP() Options {
	return func(s *server) {
		s.SetProtocal(config.HTTP)
	}
}
