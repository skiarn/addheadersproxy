package addheadersproxy

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

//Headers is used to safely store the headers when accessed concurrently by requests.
type Headers struct {
	headers []Header
	lock    *sync.RWMutex
}

//Header represents a http header with value.
type Header struct {
	Header string
	Value  string
}

type setheaders func(string, string)

//ApplyHeaders using given functon(header string, value string).
func (hs Headers) ApplyHeaders(f setheaders) {
	hs.lock.RLock()
	defer hs.lock.RUnlock()
	for _, h := range hs.headers {
		f(h.Header, h.Value)
	}
}

//String returns headers as string, nessesary when parsing flags.
func (hs *Headers) String() string {
	return fmt.Sprint(*hs)
}

//Set used when parsing flags, takes a comma-separated list, and split it into multiple header.
func (hs *Headers) Set(value string) error {
	if len(hs.headers) > 0 {
		return errors.New("headers flag already set")
	}
	for _, headers := range strings.Split(value, ",") {
		headerValue := strings.Split(headers, ":")
		if len(headerValue) != 2 {
			return fmt.Errorf("header: %v expected following format: HeaderName:Value1", headers)
		}
		hs.headers = append(h.headers, Header{Header: headerValue[0], Value: headerValue[1]})
	}
	return nil
}
