package main

import (
	"golang.org/x/net/dns/dnsmessage"
)

func extractDomain(dnsPayload []byte) (string, error) {
    var p dnsmessage.Parser
    _, err := p.Start(dnsPayload)
    if err != nil {
        return "", err
    }
    q, err := p.Question()
    if err != nil {
        return "", err
    }
    return q.Name.String(), nil
}