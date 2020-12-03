package config

// MongoDB -
type MongoDB struct {
	Host     string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	Option
}

// Redis -
type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Password string `json:"password" yaml:"password"`
	Database int    `json:"database" yaml:"database"`
	TTL      int64  `json:"ttl" yaml:"ttl"`
	Option
}

// Option -
type Option struct {
	SSL          bool  `json:"ssl" yaml:"ssl"`
	PoolSize     int64 `json:"poolSize" yaml:"poolSize"`
	MaxRetries   int64 `json:"maxRetries" yaml:"maxReties"`
	MaxIdelConns int64 `json:"maxIdelConns" yaml:"maxIdelConns"`
	MinIdelConns int64 `json:"minIdelConns" yaml:"minIdelConns"`
	MaxConns     int64 `json:"maxConns" yaml:"maxConns"`
	MinConns     int64 `json:"minConns" yaml:"minConns"`
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

// Info -
type Info struct {
	AppName string
}
