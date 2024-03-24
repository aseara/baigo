package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

var (
	_listen = flag.String("listen", ":8080", "listen address")
	_dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *_listen)
	log.Fatal(http.ListenAndServe(*_listen, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.RawPath, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		http.FileServer(http.Dir(*_dir)).ServeHTTP(resp, req)
	})))
}
