package server

type Options func(*server)

func SetQueueName(queueName string) Options {
	return func(c *server) {
		c.SetQueueName(queueName)
	}
}
