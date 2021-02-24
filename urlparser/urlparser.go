package urlparser

import (
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

//URL embeds net/url and adds extra fields ontop
type URL struct {
	Subdomain, Domain, TLD string
	ICANN                  bool
}

//Parse mirrors net/url.Parse except instead it returns
//a tld.URL, which contains extra fields.
func Parse(s string) (*URL, error) {
	dom, err := extractHost(s)
	if err != nil {
		return nil, err
	}

	//etld+1
	etld1, err := publicsuffix.EffectiveTLDPlusOne(dom)
	_, icann := publicsuffix.PublicSuffix(strings.ToLower(dom))
	if err != nil {
		return nil, err
	}

	//convert to domain name, and tld
	i := strings.Index(etld1, ".")
	domName := etld1[0:i]
	tld := etld1[i+1:]

	//and subdomain
	sub := ""
	if rest := strings.TrimSuffix(dom, "."+etld1); rest != dom {
		sub = rest
	}

	return &URL{
		Subdomain: sub,
		Domain:    domName,
		TLD:       tld,
		ICANN:     icann,
	}, nil
}

func extractHost(s string) (string, error) {
	url, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	if url.Host == "" {
		s = "http://" + s
		url, err = url.Parse(s)
		if err != nil {
			return "", err
		}
	}

	dom, _ := domainPort(url.Host)
	return dom, nil
}

func domainPort(host string) (string, string) {
	for i := len(host) - 1; i >= 0; i-- {
		if host[i] == ':' {
			return host[:i], host[i+1:]
		} else if host[i] < '0' || host[i] > '9' {
			return host, ""
		}
	}

	return host, ""
}
