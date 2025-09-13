package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 1053,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("UDP server listening on port 1053")

	buf := make([]byte, 4096)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error reading from UDP:", err)
			continue
		}

		header := string(buf[:8])
		dnsPayload := buf[8:n]

		ip, err := selectIP(header)
		if err != nil {
			fmt.Printf("Header %s: IP selection error: %v\n", header, err)
			continue
		}

		fmt.Printf("Received packet from %s, Header: %s, DNS payload len: %d, Selected IP: %s\n",
			clientAddr.String(), header, len(dnsPayload), ip)
	}
}
