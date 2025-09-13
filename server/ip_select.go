package main

import (
	"fmt"
	"strconv"
)

var IPPool = []string{
	"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5",
	"192.168.1.6", "192.168.1.7", "192.168.1.8", "192.168.1.9", "192.168.1.10",
	"192.168.1.11", "192.168.1.12", "192.168.1.13", "192.168.1.14", "192.168.1.15",
}

const ipCount = 5

func getRow(hour int) int {
	switch {
		case hour >= 4 && hour <= 11:
			return 0
		case hour >= 12 && hour <= 19:
			return 1
		default:
			return 2
	}
}

func selectIP(header string) (string,error) {
	if len(header) != 8 {
		return "", fmt.Errorf("invalid header length: %s", header)
	}

	hour, err := strconv.Atoi(header[:2])
	if err != nil {
		return "", fmt.Errorf("invalid hour in header: %v", err)
	}

	id, err := strconv.Atoi(header[6:8])
	if err != nil {
		return "", fmt.Errorf("invalid ID in header: %v", err)
	}

	startidx := getRow(hour) * ipCount 
	ipidx := startidx + (id % ipCount)
	return IPPool[ipidx], nil
}