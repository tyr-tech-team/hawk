// Package network provides network ﳑ
package network

import (
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// -
const (
	IPIFY        = "https://api.ipify.org?format=json"
	IPIFCONFIGME = "https://ifconfig.me/ip"
)

// LocalIP - 取得目前本地IP
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

// HostIP - 取得目前本機IP
func HostIP() string {
	con, _ := net.Dial("udp", "8.8.8.8:80")
	defer con.Close()
	ip := con.LocalAddr().String()
	xip := strings.Split(ip, ":")
	return xip[0]
}

// PublishIP - 取的目前對外IP
func PublishIP() string {
	resp, err := http.Get(IPIFCONFIGME)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
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
