package redis

import (
	"context"
	"fmt"

	goredis "github.com/go-redis/redis/v8"
	"github.com/tyr-tech-team/hawk/config"
	"github.com/tyr-tech-team/hawk/status"
)

// NewDial -
func NewDial(c config.Redis) (*goredis.Client, error) {

	ctx := context.TODO()
	opts := goredis.Options{
		Addr:         c.Host,
		Password:     c.Password,
		DB:           c.Database,
		PoolSize:     int(config.DefaultMaxPoolSize),
		MaxRetries:   config.DefaultMaxRetries,
		MinIdleConns: int(config.DefaultMinPoolSize),
	}

	if c.MaxPoolSize > 0 {
		opts.PoolSize = int(c.MaxPoolSize)
	}

	if c.MaxRetries > 0 {
		opts.MaxRetries = int(c.MaxRetries)
	}

	if c.MinIdelConns > 0 {
		opts.MinIdleConns = int(c.MinIdelConns)
	}

	client := goredis.NewClient(&opts)
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, status.ConnectFailed.WithDetail(fmt.Sprintf("連結 Redis 失敗: %s ", err))
	}

	return client, nil
}
