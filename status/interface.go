package status

// Status -
type Status interface {
	Err() error
	Error() string
	String() string
	Status() Status
	Detail() []string
	Code() string
	WithDetail(detail ...string) Status
	SetServiceCode(serviceCode ServiceCode) Status
}
