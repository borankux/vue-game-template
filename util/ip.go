package util

import (
	"net"
	"strconv"
	"strings"
)

func GetIPAndPort(addr net.Addr) (ip string, port int) {
	split := strings.Split(addr.String(), ":")
	port, _ = strconv.Atoi(split[1])
	ip = split[0]
	return ip, port
}
