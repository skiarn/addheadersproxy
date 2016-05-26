# go-addheadersproxy

## Manually
### How to get started ###

* Install go - https://golang.org/dl/
* git clone https://github.com/skiarn/go-addheadersproxy

## Example use.

```
import(
  )

  package main

  import (
    "github.com/skiarn/go-addheadersproxy"
  	"log"
  	"net/http"
  )

  func main() {
  	NewTestServer(8000)
  	proxy := ReverseProxy("http://localhost:8000")
  	mux := http.NewServeMux()
  	mux.Handle("/", proxy)
  	log.Fatal(http.ListenAndServe(":8090", mux))
  }


```

## Run application
./jenkins-plugins
