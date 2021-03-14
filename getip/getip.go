package getip

import (
	"fmt"
	"net"
)

func GetIP(hostName string) (string, error) {
	ips, err := net.LookupIP(hostName)
	if err != nil {
		return "", fmt.Errorf("faild to lookup IP against host; %s; %w", hostName, err)
	}
	for _, ip := range ips {
		if ip.To4() != nil {
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("failed to get IP")
}
