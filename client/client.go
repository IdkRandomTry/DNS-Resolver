package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Take pcap file as input
	pcapFile := flag.String("f", "1.pcap", "Path to the PCAP file")
	serverAddr := flag.String("s", "127.0.0.1:1053", "Server UDP address")
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", *serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Open file
	handle, err := pcap.OpenOffline(*pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var id int = 0

	// Set up packet source
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Iterate over packets
	for packet := range packetSource.Packets() {
		// Extract DNS layer
		dnsLayer := packet.Layer(layers.LayerTypeDNS)
		// Skip if no DNS layer
		if dnsLayer == nil {
			continue
		}
		dns := dnsLayer.(*layers.DNS)
		// Skip if query response
		if dns.QR {
			continue
		}

		// Get UTC timestamp
		ts := time.Now().UTC()
		tsStr := ts.Format("150405")
		idStr := fmt.Sprintf("%02d", id%100)
		header := []byte(tsStr + idStr) // 8-byte header

		dnsOriginal := dnsLayer.LayerContents()
		dnsModified := append(header, dnsOriginal...)

		// Send modified DNS query over UDP
		_, err := conn.Write(dnsModified)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sent DNS query with ID %d at %s\n", id, tsStr)

		id++
	}

	// Wait for user input before closing
	fmt.Println("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
