package status

import (
	"encoding/json"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"
)

type status struct {
	gst  *gs.Status
	body *body
}

// NewStatus - 新增狀態碼
func NewStatus(levelCode LevelCode, serviceCode ServiceCode, grpcCode GRPCCode, actionCode ActionCode, msg string) Status {
	b := &body{
		lCode:   levelCode,
		sCode:   serviceCode,
		gCode:   grpcCode,
		aCode:   actionCode,
		Message: msg,
	}

	s := new(status)
	s.body = b
	s.gst = gs.New(codes.Code(grpcCode), b.Marshal())

	return s
}

func (s *status) Error() string {
	return s.gst.Message()
}

func (s *status) String() string {
	return fmt.Sprintf("[%s] %s", s.body.Code, s.body.Message)
}

func ConvertStatus(err error) (Status, error) {

	if err == nil {
		return NoError, nil
	}

	gsError, ok := gs.FromError(err)
	if !ok {
		return nil, errors.New("convert error failed")
	}

	b := new(body)
	if err := json.Unmarshal([]byte(gsError.Message()), b); err != nil {
		return nil, errors.New("can't parse the error message")
	}

	s := &status{
		gst:  gsError,
		body: b,
	}

	return s, nil
}
