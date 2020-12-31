package srv

// Register -
type Register interface {
	SetRegisterConfig(s ServiceRegisterConfig)
	Register() error
	Deregister() error
	Close()
}

// ServiceRegisterConfig -
type ServiceRegisterConfig struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`
	Port     int      `json:"port"`
	Address  string   `json:"address"`
	Traefik  bool     `json:"traefik"`
	Protocol Protocol `json:"protocol"`
}
