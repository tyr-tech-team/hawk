package config

// Config -
type Config struct {
	Info    Info    `json:"info" yaml:"info"`
	Service Service `json:"service" yaml:"service"`
	Mongo   MongoDB `json:"mongo" yaml:"mongo"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	Nats    Nats    `json:"nats" yaml:"nats"`
}

// MongoDB -
type MongoDB struct {
	Hosts    []string `json:"host" yaml:"host"`
	User     string   `json:"user" yaml:"user"`
	Password string   `json:"password" yaml:"password"`
	Database string   `json:"database" yaml:"database"`
	Option
}

// Redis -
type Redis struct {
	Hosts    []string `json:"host" yaml:"host"`
	Password string   `json:"password" yaml:"password"`
	Database int      `json:"database" yaml:"database"`
	TTL      int64    `json:"ttl" yaml:"ttl"`
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
	Hosts []string `json:"host" yaml:"host"`
}

// Service -
type Service map[string]string

// Log -
type Log struct {
	Level string `json:"level" yaml:"level"`
}

// Info -
type Info struct {
	Name    string `json:"Name" yaml:"Name"`
	Address string `json:"address" yaml:"address"`
	Port    int    `json:"port" yaml:"port"`
}
