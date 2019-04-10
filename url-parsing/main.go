package main

import "fmt"
import "net"
import "net/url"

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	p := fmt.Println

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p(u.Scheme)
	p(u.User)
	p(u.User.Username())
	pw, _ := u.User.Password()
	p(pw)

	p(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	p(host)
	p(port)
	p(u.Path)
	p(u.Fragment)
	p(u.RawQuery)
	q, _ := url.ParseQuery(u.RawQuery)
	p(q)
	p(q["k"][0])

}
