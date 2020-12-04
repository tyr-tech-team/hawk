package tools

import (
	"net"
	"strconv"
	"strings"

	"github.com/parnurzeal/gorequest"
)

const (
	IPIFY = "https://api.ipify.org?format=json"
)

// LocalIP -
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, v := range addrs {
		if ip, ok := v.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}
	return ""
}

// HostIP -
func HostIP() string {
	con, _ := net.Dial("udp", "8.8.8.8:80")
	defer con.Close()
	ip := con.LocalAddr().String()
	xip := strings.Split(ip, ":")
	return xip[0]
}

// PublishIP -
func PublishIP() string {
	r := gorequest.New()
	x := struct {
		IP string `json:"ip"`
	}{}

	_, _, errs := r.Get(IPIFY).EndStruct(&x)
	for _, err := range errs {
		if err != nil {
			return ""
		}
	}
	return x.IP
}

// GetFreePort - 取得目前可以使用的 Port
func GetFreePort() (int, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	addr := listener.Addr().String()
	_, portString, err := net.SplitHostPort(addr)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(portString)
}
