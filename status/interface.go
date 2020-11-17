package status

// Status -
type Status interface {
	Err() error
	Error() string
	String() string
	//
	Code() string
	Detail() []string
	Status() Status
	Message() string
	EMessage() string
	//
	WithDetail(detail ...string) Status
	SetServiceCode(serviceCode ServiceCode) Status
}
