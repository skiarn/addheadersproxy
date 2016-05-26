# go-addheadersproxy

## Manually
### How to get started ###

* Install go - https://golang.org/dl/
* git clone https://github.com/skiarn/go-addheadersproxy

## Example use.

```
package main

import (
	"log"
	"net/http"

	"github.com/skiarn/addheadersproxy"
	"github.com/skiarn/addheadersproxy/test"
)

func main() {
	test.NewTestServer(8000)
	proxy := addheadersproxy.ReverseProxy("http://localhost:8000")
	mux := http.NewServeMux()
	mux.Handle("/", proxy)
	log.Fatal(http.ListenAndServe(":8090", mux))
}

```
