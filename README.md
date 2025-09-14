# DNS Resolver Project

## Overview

This project implements a **client-server DNS resolver** simulation in Go:

Selecting PCAP: (Siddhesh Umarjee: 23110347  + Shounak Ranade: 23110304 )%10 = 1.pcap

- **Client**: Reads a PCAP file, filters **DNS query packets**, prepends a **custom 8-byte header (HHMMSSID)**, and sends them to the server over **UDP**.
- **Server**: Receives packets, extracts the custom header, applies **time-based IP selection rules** for load balancing, and logs the resolved IP along with the original header.

**Custom Header Format (HHMMSSID)**:  
- HH → hour in UTC (24-hour format)  
- MM → minutes in UTC  
- SS → seconds in UTC  
- ID → sequential ID of DNS query (00, 01, …)  

**Server IP Selection**:  
- Pool of 15 IPs split into 3 time slots: morning (4:00–11:59), afternoon (12:00–19:59), night (20:00–3:59).  
- Selection formula: `final_index = ip_slot_start + (ID % 5)`

---

## Folder Structure
```
dns_resolver/
- client/
  - client.go       # Client code
  - 1.pcap          # PCAP file
- server/
  - server.go       # UDP server
  - ip_select.go    # IP selection logic
- .gitignore        # Recommended ignore rules
- go.mod            # Go module file
```

---

## Prerequisites

- **Go 1.18+** installed  
- Run following command to install dependencies
```
    go mod tidy
```

---

## Running the Server

1. Open a terminal in the `server/` folder:
```
    go run server.go
```
2. The server listens on **UDP port 1053**.  
3. You should see:
```
    UDP server listening on port 1053
```
---

## Running the Client

1. Open a separate terminal in the `client/` folder:
```
    go run client.go -f input.pcap -s 127.0.0.1:1053    
```
run without flags for default test (1.pcap)

- `-f` → Path to the PCAP file  
- `-s` → Server IP and UDP port  

2. You should see output like:
    Sent DNS query with ID 0 at 105827
    Sent DNS query with ID 1 at 105827
    Press 'Enter' to exit...
---

## Server Output Example
       
    Received packet from 127.0.0.1:59719, Header: 11100400, DNS payload len: 30, Selected IP: 192.168.1.1        
    Received packet from 127.0.0.1:59719, Header: 11100401, DNS payload len: 35, Selected IP: 192.168.1.2  
    ...
    Received packet from 127.0.0.1:57919, Header: 12505100, DNS payload len: 30, Selected IP: 192.168.1.6

---

## Notes

- **Custom header is generated using the UTC timestamp of packet processing**, not the original PCAP timestamp.  
- Only **DNS query packets** are sent to the server.  
- IP selection is **deterministic** based on the header and time slot. 
