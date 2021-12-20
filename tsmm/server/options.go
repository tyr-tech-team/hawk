package server

// Options -
type Options func(*server)

// SetQueueName - 設定Queue group name option
func SetQueueName(queueName string) Options {
	return func(c *server) {
		c.SetQueueName(queueName)
	}
}
