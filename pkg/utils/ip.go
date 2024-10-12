package utils

import (
	"fmt"
	"net"
)

func GetLocalIp() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		err = fmt.Errorf("failed to establish conn: %s", err)
		return "", err
	}

	lAddr := conn.LocalAddr().(*net.UDPAddr)
	ipAddr := lAddr.IP.String()

	return ipAddr, nil
}
