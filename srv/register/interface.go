package register

// Register -
type Register interface {
	Register() error
	Deregister() error
}
