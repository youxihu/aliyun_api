package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetPublicIP(stripCidr bool) (string, error) {
	resp, err := http.Get("https://ifconfig.me/ip")
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	ipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ip := strings.TrimSpace(string(ipBytes))
	if ip == "" {
		return "", fmt.Errorf("无法获取公网 IP")
	}

	if stripCidr {
		return ip, nil
	}

	return ip + "/32", nil
}
