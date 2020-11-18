package status

import (
	"encoding/json"
	"fmt"

	"google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"
)

type status struct {
	gst  *gs.Status
	body *body
}

// NewStatus - 新增狀態碼
func NewStatus(levelCode LevelCode, serviceCode ServiceCode, grpcCode GRPCCode, actionCode ActionCode, msg string, emsg ...string) Status {
	b := &body{
		lCode:   levelCode,
		sCode:   serviceCode,
		gCode:   grpcCode,
		aCode:   actionCode,
		Message: msg,
		Details: []string{},
	}

	if len(emsg) > 0 {
		b.EMessage = emsg[0]
	}

	s := new(status)
	s.body = b
	s.gst = gs.New(codes.Code(grpcCode), b.Marshal())

	return s
}

// ConvertStatus -
func ConvertStatus(err error) Status {
	if err == nil {
		return NoError
	}

	gsError, ok := gs.FromError(err)
	if !ok {
		fmt.Println(gsError.Err(), ok)
		return UnKnownError
	}

	s := new(status)

	s.gst = gsError
	b := new(body)
	json.Unmarshal([]byte(gsError.Message()), b)
	b.ParseCode()

	s.body = b

	return s
}

// Equal -
func Equal(s1 Status, s2 Status) bool {
	switch {
	case s1.Code() != s2.Code():
	case s1.Message() != s2.Message():
	default:
		return true
	}
	return false
}

func (s status) Error() string {
	return s.gst.Err().Error()
}

func (s status) Err() error {
	return s.gst.Err()
}

func (s status) String() string {
	return fmt.Sprintf("code[%s] message: %s", s.body.Code, s.body.Message)
}

func (s status) WithDetail(detail ...string) Status {
	newbody := copyBody(s.body)
	newbody.Details = append(newbody.Details, detail...)
	s.body = newbody
	return s.marshal()
}

func (s status) SetServiceCode(serviceCode ServiceCode) Status {
	newbody := copyBody(s.body)
	newbody.SetSCode(serviceCode)
	s.body = newbody
	return s.marshal()
}

func (s status) Detail() []string {
	return s.body.Details
}

func (s status) Code() string {
	return s.body.Code
}

func (s status) Message() string {
	return s.body.Message
}

func (s status) EMessage() string {
	return s.body.EMessage
}

func (s status) marshal() Status {
	s.gst = gs.New(s.gst.Code(), s.body.Marshal())
	return s
}

func (s status) Status() Status {
	gsError, ok := gs.FromError(s.Err())
	if !ok {
		fmt.Println(gsError.Err(), ok)
		return UnKnownError
	}
	s.gst = gsError
	b := new(body)
	json.Unmarshal([]byte(gsError.Message()), b)
	b.ParseCode()

	s.body = b

	return s
}
