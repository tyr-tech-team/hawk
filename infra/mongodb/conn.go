package mongodb

import (
	"context"
	"crypto/tls"
	"log"
	"strings"

	"github.com/tyr-tech-team/hawk/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewDial -
func NewDial(c config.Mongo) (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultTimeout)
	defer cancel()

	opts := options.Client()
	opts.SetAppName(c.Name)

	opts.SetHosts(strings.Split(c.Host, ","))

	if c.User != "" && c.Password != "" {
		opts.SetAuth(options.Credential{
			Username:   c.User,
			Password:   c.Password,
			AuthSource: c.Database,
		})
	}
	if c.ReplicaSet != "" {
		opts.ReplicaSet = &c.ReplicaSet
	}

	if c.SSL {
		opts.SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	}

	if c.MaxPoolSize > 0 {
		opts.SetMaxPoolSize(c.MaxPoolSize)
	} else {
		opts.SetMaxPoolSize(config.DefaultMaxPoolSize)
	}

	if c.MinPoolSize > 0 {
		opts.SetMinPoolSize(c.MinPoolSize)
	} else {
		opts.SetMinPoolSize(config.DefaultMinPoolSize)
	}

	if c.MaxConnIdleTime > 0 {
		opts.SetMaxConnIdleTime(c.MaxConnIdleTime)
	} else {
		opts.SetMaxConnIdleTime(config.DefaultMaxConnIdelTime)
	}

	if c.HeartbeatInterval > 0 {
		opts.SetHeartbeatInterval(c.HeartbeatInterval)
	} else {
		opts.SetHeartbeatInterval(config.DefaultHeartbeatInterval)
	}

	opts.SetDirect(c.Direct)

	// Connect to MongoDB
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		log.Fatal("connect failed in mongo", err)
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("ping failed in mongo", err)
		return nil, err
	}

	return client, nil
}
