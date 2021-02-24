package dgrep

import "regexp"

// Ref: https://stackoverflow.com/questions/2180465/can-domain-name-subdomains-have-an-underscore-in-it
var domainRegex = regexp.MustCompile(`[-a-zA-Z0-9._]{2,256}\.[a-z]{2,}`)

// ExtractDomains will return all domains found in the input file
func ExtractDomains(inputStrs []string) []string {
	return extractWithRegex(inputStrs, domainRegex)
}
