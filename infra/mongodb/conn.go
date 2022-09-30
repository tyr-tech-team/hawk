// Package mongodb provides mongodb ﳑ
package mongodb

import (
	"context"
	"crypto/tls"
	"log"
	"strings"
	"time"

	"github.com/tyr-tech-team/hawk/config"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Options -
type Options struct{}

// Option -
type Option func(o *options.ClientOptions)

// NewDial -
func NewDial(c config.Mongo, option ...Option) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultTimeout)
	defer cancel()

	// new options
	opts := newOptions()

	// set default -
	defaultOption := setDefault(opts, c)

	// append option
	option = append(option, defaultOption...)
	for _, o := range option {
		o(opts)
	}

	// Connect to MongoDB
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	// connect
	if err := client.Connect(ctx); err != nil {
		log.Fatal("connect failed in mongo", err)
		return nil, err
	}

	// ping connect
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("ping failed in mongo", err)
		return nil, err
	}

	return client, nil
}

// new option
func newOptions() *options.ClientOptions {
	return options.Client()
}

// setDefault - set connection options
func setDefault(opts *options.ClientOptions, c config.Mongo) []Option {
	o := make([]Option, 0)

	// set default
	opts.SetAppName(c.Name)

	// hosts
	opts.SetHosts(strings.Split(c.Host, ","))

	// max pool
	opts.SetMaxPoolSize(config.DefaultMaxPoolSize)

	// min pool
	opts.SetMinPoolSize(config.DefaultMinPoolSize)

	// max ConnIdleTime
	opts.SetMaxConnIdleTime(config.DefaultMaxConnIdelTime)

	// Heartbeat Interval
	opts.SetHeartbeatInterval(config.DefaultHeartbeatInterval)

	// set value
	if c.User != "" && c.Password != "" {
		o = append(o, SetAuth(options.Credential{
			Username:   c.User,
			Password:   c.Password,
			AuthSource: c.Database,
		}))
	}

	// set ReplicaSet
	if c.ReplicaSet != "" {
		o = append(o, SetReplica(&c.ReplicaSet))
	}

	// set ssl
	if c.SSL {
		o = append(o, SetSSL())
	}

	// set MaxPoolSize
	if c.MaxPoolSize > 0 {
		o = append(o, SetMaxPoolSize(c.MaxPoolSize))
	}

	// set MinPoolSize
	if c.MinPoolSize > 0 {
		o = append(o, SetMinPoolSize(c.MinPoolSize))
	}

	// set MaxConnIdleTime
	if c.MaxConnIdleTime > 0 {
		o = append(o, SetMaxConnIdleTime(c.MaxConnIdleTime))
	}

	// set HeartbeatInterval
	if c.HeartbeatInterval > 0 {
		o = append(o, SetHeartbeatInterval(c.HeartbeatInterval))
	}

	// SetDirect
	opts.SetDirect(c.Direct)

	return o
}

// SetMonitor -
func SetMonitor(m *event.CommandMonitor) Option {
	return func(o *options.ClientOptions) {
		o.SetMonitor(m)
	}
}

// SetMaxPoolSize -
func SetMaxPoolSize(u uint64) Option {
	return func(o *options.ClientOptions) {
		o.SetMaxPoolSize(u)
	}
}

// SetMinPoolSize -
func SetMinPoolSize(u uint64) Option {
	return func(o *options.ClientOptions) {
		o.SetMinPoolSize(u)
	}
}

// SetMaxConnIdleTime -
func SetMaxConnIdleTime(t time.Duration) Option {
	return func(o *options.ClientOptions) {
		o.SetMaxConnIdleTime(t)
	}
}

// SetHeartbeatInterval -
func SetHeartbeatInterval(t time.Duration) Option {
	return func(o *options.ClientOptions) {
		o.SetHeartbeatInterval(t)
	}
}

// SetAuth -
func SetAuth(a options.Credential) Option {
	return func(o *options.ClientOptions) {
		o.SetAuth(a)
	}
}

// SetReplica -
func SetReplica(s *string) Option {
	return func(o *options.ClientOptions) {
		o.ReplicaSet = s
	}
}

// SetSSL -
func SetSSL() Option {
	return func(o *options.ClientOptions) {
		o.SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	}
}
