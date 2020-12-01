package config

// MongoDB -
type MongoDB struct {
	Host     string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	SSL      bool   `json:"ssl" yaml:"ssl"`
}

// Redis -
type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Password string `json:"password" yaml:"password"`
	Database int    `json:"database" yaml:"database"`
	TTL      int64  `json:"ttl" yaml:"ttl"`
}

// Option -
type Option struct {
	MaxIdel int64
	MinIdel int64
	MaxConn int64
	MinConn int64
}

// Nats -
type Nats struct {
}

// Service -
type Service struct {
}

// Log -
type Log struct {
	Level string `json:"level"`
}
