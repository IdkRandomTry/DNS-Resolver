package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 1053,
		IP:   net.ParseIP("0.0.0.0"),  //not 127.0.0.1 so that server listens on all interfaces
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("UDP server listening on port 1053")
	fmt.Println("Custom Header\t | Domain Name\t\t | Resolved IP")
	fmt.Println("-----------------|-----------------------|----------------")

	buf := make([]byte, 4096)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error reading from UDP:", err)
			continue
		}

		header := string(buf[:8])
		domain, err := extractDomain(buf[8:n])
		
		if err != nil {
			fmt.Printf("Header %s: Domain extraction error: %v\n", header, err)
			continue
		}

		ip, err := selectIP(header)
		if err != nil {
			fmt.Printf("Header %s: IP selection error: %v\n", header, err)
			continue
		}

		fmt.Printf("%s \t | %-21s | %s  \n", header, domain, ip)
	}
}
