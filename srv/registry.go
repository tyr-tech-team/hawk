package srv

import "github.com/tyr-tech-team/hawk/config"

// Register -
type Register interface {
	SetRegisterConfig(s config.ServiceRegister)
	Register() error
	Deregister() error
	Close()
}
