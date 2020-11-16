package status

// Status -
type Status interface {
	Err() error
	Error() string
	String() string
	Detail() []string
	Status() Status
	Detail() []string
	Code() string
	WithDetail(detail ...string) Status
	SetServiceCode(serviceCode ServiceCode) Status
}
