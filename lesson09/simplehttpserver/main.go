package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		_, err := rw.Write([]byte("Hello, world"))
		if err != nil {
			logrus.Warn(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
