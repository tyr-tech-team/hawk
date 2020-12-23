package config

import "time"

// Config - 基礎配置檔
type Config struct {
	Info    Info    `json:"info" yaml:"info"`
	Service Service `json:"service" yaml:"service"`
	Mongo   MongoDB `json:"mongo" yaml:"mongo"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	Nats    Nats    `json:"nats" yaml:"nats"`
}

// MongoDB - Mongo資料庫配置檔
type MongoDB struct {
	Hosts    string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	DatabaseOption
}

// Redis - Redis 資料庫配置
type Redis struct {
	Hosts    string `json:"host" yaml:"host"`
	Password string `json:"password" yaml:"password"`
	Database int    `json:"database" yaml:"database"`
	TTL      int64  `json:"ttl" yaml:"ttl"`
	DatabaseOption
}

// DatabaseOption - 資料庫額外參數
type DatabaseOption struct {
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
	Hosts string `json:"host" yaml:"host"`
}

// Service - 服務名稱配置
type Service map[string]string

// Log - 紀錄配置檔
type Log struct {
	Level string `json:"level" yaml:"level"`
}

// Info - 服務基本資訊
type Info struct {
	Name   string `json:"Name" yaml:"Name"`
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Consul string `json:"consul" yaml:"consul"`
	Mode   string `json:"mode"`
}

// ServiceRegister - 服務註冊使用配置檔
type ServiceRegister struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`
	Port     int      `json:"port"`
	Address  string   `json:"address"`
	Traefik  bool     `json:"traefik"`
	Protocol Protocol `json:"protocol"`
}

// ChangeSet  -
type ChangeSet struct {
	Data      []byte
	Checksum  string
	Timestamp time.Time
}
