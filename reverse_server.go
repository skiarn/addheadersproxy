package addheadersproxy

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"
)

var headersFlag Headers

func init() {
	flag.Var(&headersFlag, "headers", "comma-separated list of header to be applied use following format header1:value1,header2:value2, example: -headers=header1:value1,Cool-try-Hard:true")
}

//ReverseProxy takes a target URI ex. http://localhost:8000 and directs trafic there.
func ReverseProxy(target string) *httputil.ReverseProxy {
	flag.Parse()
	headersFlag.lock = new(sync.RWMutex)
	fmt.Printf("Headers:  %v \n", headersFlag.headers)
	targetParts := strings.Split(target, "://")
	if len(targetParts) != 2 {
		log.Fatalf("Invalid target URI")
	}
	scheme := targetParts[0]
	host := targetParts[1]
	director := func(req *http.Request) {
		//TODO: X-Forwarded-For a de facto standard for identifying the originating IP address of a client connecting to a web server through an HTTP proxy or load balancer.
		// Example: X-Forwarded-For: client1, proxy1, proxy2, X-Forwarded-For: 129.78.138.66, 129.78.64.103
		// TODO: X-Forwarded-Host a de facto standard for identifying the original host requested by the client in the Host HTTP request header, since the host name and/or port of the reverse proxy (load balancer) may differ from the origin server handling the request.
		// Example: X-Forwarded-Host: en.wikipedia.org:80, X-Forwarded-Host: en.wikipedia.org
		req.URL.Scheme = scheme
		req.URL.Host = host
	}
	dial := func(network, addr string) (net.Conn, error) {
		log.Printf("Dial to addr: %v, netowrk: %v", addr, network)
		return net.Dial(network, addr)
	}
	transport := &customTransport{&http.Transport{Dial: dial}}
	res := &httputil.ReverseProxy{Director: director, Transport: transport}
	return res
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	headersFlag.ApplyHeaders(req.Header.Add)
	res, err := t.Transport.RoundTrip(req)
	return res, err
}

type customTransport struct {
	*http.Transport
}
