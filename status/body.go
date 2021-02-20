package status

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type body struct {
	lCode    LevelCode
	sCode    ServiceCode
	gCode    GRPCCode
	dCode    DescCode
	Code     string   `json:"code"`
	Message  string   `json:"message,omitempty"`
	EMessage string   `json:"emessage,omitempty"`
	Details  []string `json:"details,omitempty"`
}

func (b *body) fromatCode() {
	b.Code = fmt.Sprintf("%s-%s-%s-%s", b.lCode.String(), b.sCode.String(), b.gCode.String(), b.dCode.String())
}

// Marshal -
func (b *body) Marshal() string {
	b.fromatCode()

	str, err := json.Marshal(b)
	if err != nil {
		log.Panicf("marshal json failed with body[%+v]", b)
		return ""
	}
	return string(str)
}

func (b *body) ParseCode() error {
	codeList := strings.Split(b.Code, "-")

	if len(codeList) != 4 {
		return errors.New("wrong code fromat")
	}

	b.lCode = LevelCode(convertToInt64(codeList[0]))
	b.sCode = ServiceCode(convertToInt64(codeList[1]))
	b.gCode = GRPCCode(convertToInt64(codeList[2]))
	b.dCode = DescCode(convertToInt64(codeList[3]))

	return nil
}

func (b *body) SetSCode(serviceCode ServiceCode) {
	b.sCode = serviceCode
}

func convertToInt64(s string) int64 {
	tmp, err := strconv.ParseInt(s, 32, 10)
	if err != nil {
		log.Panicf("convert string to int64 failed [%v]", err)
		return 0
	}

	return tmp
}

func copyBody(b *body) *body {
	x := &body{
		lCode:    b.lCode,
		sCode:    b.sCode,
		gCode:    b.gCode,
		dCode:    b.dCode,
		Code:     b.Code,
		Message:  b.Message,
		EMessage: b.EMessage,
		Details:  b.Details,
	}
	return x
}
