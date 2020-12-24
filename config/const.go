package config

import "time"

// Protocol -
type Protocol string

//
const (
	GRPC Protocol = "grpc"
	HTTP Protocol = "http"
)

// 常數
const (
	// DefaultTimeout - 初始超時時間
	DefaultTimeout = 10 * time.Second
	// DefaultPageSize - 初始每頁數量
	DefaultPageSize = 10
)

const (
	// DefaultMaxConnIdelTime - 初始每條通道閒置時間
	DefaultMaxConnIdelTime = time.Second * 300
	// DefaultHeartbeatInterval -
	DefaultHeartbeatInterval = time.Second * 10
	// DefaultMaxPoolSize - 初始最大連接池數量
	DefaultMaxPoolSize uint64 = 5
	// DefaultMinPoolSize - 初始最小連接池數量
	DefaultMinPoolSize uint64 = 3
	// DefaultMaxRetries - 最大重試次數
	DefaultMaxRetries int = 3
)
