package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	urlinfo, _ := url.Parse("postgres://packie:1232323@host.com:5432/path?k=v#f")

	fmt.Printf("%-30s = %s\n", "urlinfo.Scheme", urlinfo.Scheme)
	fmt.Printf("%-30s = %s\n", "urlinfo.User", urlinfo.User)
	fmt.Printf("%-30s = %s\n", "urlinfo.User.Username()", urlinfo.User.Username())
	fmt.Printf("%-30s = %s\n", "urlinfo.Host", urlinfo.Host)
	fmt.Printf("%-30s = %s\n", "urlinfo.Path", urlinfo.Path)
	fmt.Printf("%-30s = %s\n", "urlinfo.Fragment", urlinfo.Fragment)
	fmt.Printf("%-30s = %s\n", "urlinfo.RawQuery", urlinfo.RawQuery)

	p, _ := urlinfo.User.Password()
	fmt.Printf("%-30s = %s\n", "urlinfo.User.Password()", p)

	host, port, _ := net.SplitHostPort(urlinfo.Host)
	fmt.Printf("%-30s = %s\n", "host", host)
	fmt.Printf("%-30s = %s\n", "port", port)

	m, _ := url.ParseQuery(urlinfo.RawQuery)
	fmt.Printf("%-30s = %s\n", "m", m)
	fmt.Printf("%-30s = %s\n", "m[\"k\"][0]", m["k"][0])
}
