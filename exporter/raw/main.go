package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, `exporter_request_count{user="admin"} 1000`)
}

func main() {
	http.HandleFunc("/metrics", helloHandler)
	_ = http.ListenAndServe(":8050", nil)
}
