package testtarget

import (
	"io"
	"net/http"
	"strconv"
)

func NewTestServer(port int) {
	muxtest := http.NewServeMux()
	muxtest.HandleFunc("/", testServer)
	go http.ListenAndServe(":"+strconv.Itoa(port), muxtest)
}

func testServer(w http.ResponseWriter, r *http.Request) {
	headers := ""
	for k, v := range r.Header {
		for _, vv := range v {
			headers = headers + k + ":" + vv + "\n"
		}
	}
	io.WriteString(w, "Im a test server!\n I recive headers:\n"+headers)
}
