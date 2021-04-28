package config

import "time"

// DefaultConfig - 基礎配置檔
type DefaultConfig struct {
	Info    Info    `json:"info" yaml:"info"`
	Service Service `json:"service" yaml:"service"`
	Mongo   Mongo   `json:"mongo" yaml:"mongo"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	Nats    Nats    `json:"nats" yaml:"nats"`
	Log     Log     `json:"log" yaml:"log"`
}

// Mongo - Mongo資料庫配置檔
type Mongo struct {
	Name       string `json:"name" yaml:"name"`
	Host       string `json:"host" yaml:"host"`
	User       string `json:"user" yaml:"user"`
	Password   string `json:"password" yaml:"password"`
	Database   string `json:"database" yaml:"database"`
	ReplicaSet string `json:"replicaSet" yaml:"replicaSet"`
	DatabaseOption
}

// Redis - Redis 資料庫配置
type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Password string `json:"password" yaml:"password"`
	Database int    `json:"database" yaml:"database"`
	TTL      int64  `json:"ttl" yaml:"ttl"`
	DatabaseOption
}

// DatabaseOption - 資料庫額外參數
type DatabaseOption struct {
	SSL               bool          `json:"ssl" yaml:"ssl"`
	MaxPoolSize       uint64        `json:"maxPoolSize" yaml:"maxPoolSize"`
	MinPoolSize       uint64        `json:"minPoolSize" yaml:"minPoolSize"`
	MaxRetries        int64         `json:"maxRetries" yaml:"maxReties"`
	MaxIdelConns      int64         `json:"maxIdelConns" yaml:"maxIdelConns"`
	MinIdelConns      uint64        `json:"minIdelConns" yaml:"minIdelConns"`
	MaxConns          int64         `json:"maxConns" yaml:"maxConns"`
	MinConns          int64         `json:"minConns" yaml:"minConns"`
	MaxConnIdleTime   time.Duration `json:"maxConnIdleTime" yaml:"maxConnIdleTime"`
	HeartbeatInterval time.Duration `json:"heartbeatInterval" yaml:"heartbeatInterval"`
	Direct            bool          `json:"direct" yaml:"direct"`
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
	// Name - 服務名稱
	Name string `json:"Name" yaml:"Name"`
	// RemoteHost -  遠端Config 位置
	RemoteHost string `json:"remoteHost" yaml:"remoteHost"`
	// Port - 監聽通訊埠
	Port string `json:"port" yaml:"port"`
	// Host -  監聽網路介面
	Host string `json:"host" yaml:"host"`
	// Mod - 啟動模式
	Mode string `json:"mode"`
	// Version -  版本號
	Version string `json:"version" yaml:"version"`
	// Git commit
	Commit string `json:"commit" yaml:"commit"`
	// Build - 建置時間
	Build string `json:"build" yaml:"build"`
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

// Operator - 操作者資訊
type Operator struct {
	// Name - 更新者姓名
	Name string `json:"name"`
	// Account - 帳號
	Account string `json:"account"`
	// Identifier - 身份類型
	Identifier int32 `json:"identifier"`
	// Time - 操作時間
	Time time.Time `json:"time"`
}
