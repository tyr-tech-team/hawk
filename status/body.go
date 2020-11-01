package status

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

type body struct {
	lCode   LevelCode
	sCode   ServiceCode
	gCode   GRPCCode
	aCode   ActionCode
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (b *body) fromatCode() {
	b.Code = fmt.Sprintf("%s-%s-%s-%s", b.lCode.String(), b.sCode.String(), b.gCode.String(), b.aCode.String())
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

	return nil
}
