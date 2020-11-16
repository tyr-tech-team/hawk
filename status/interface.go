package status

// Status -
type Status interface {
	Err() error
	Error() string
	String() string
	Code() string
	Detail() []string
	Status() Status
	WithDetail(detail ...string) Status
	SetServiceCode(serviceCode ServiceCode) Status
}
