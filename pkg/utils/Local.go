package utils

import (
	"errors"
	"net"
)

// GetLocalIP Obtain local IP address.
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ip, ok := addr.(*net.IPNet)
		if ok {
			return ip.IP.String(), nil
		}
	}
	return "", errors.New("failed to obtain the local IP address")
}
