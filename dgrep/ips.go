package dgrep

import (
	"net"
	"regexp"

	"github.com/audiolion/ipip"
)

// Ref: https://stackoverflow.com/questions/2180465/can-domain-name-subdomains-have-an-underscore-in-it
var ipRegex = regexp.MustCompile(`\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b`)

// ExtractIPs will return all IP addresses found in the input file
func ExtractIPs(inputStrs []string) []net.IP {
	addrStrs := extractWithRegex(inputStrs, ipRegex)
	addrs := make([]net.IP, 0, len(addrStrs))

	for _, addrStr := range addrStrs {
		addr := net.ParseIP(addrStr)
		if addr != nil {
			addrs = append(addrs, addr)
		}
	}

	return addrs
}

// GetPublicIPs will return only non-private IP addresses
// as defined in RFC 1918.
func GetPublicIPs(addrs []net.IP) []net.IP {
	var publicAddrs []net.IP

	for _, addr := range addrs {
		if ipip.IsPrivate(addr) || addr.IsLoopback() {
			continue
		}

		publicAddrs = append(publicAddrs, addr)
	}

	return publicAddrs
}
