package middleware

import (
	"log"
	"mime"
	"net/http"
)

const errNoMediaTypeStr = "mime: no media type"

// Logging log
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("recv a %s request from %s", req.Method, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

// Validating validate
func Validating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")
		mediatype, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			if err.Error() != errNoMediaTypeStr {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			mediatype = "application/json"
		}
		if mediatype != "application/json" {
			http.Error(w, "invalid Content-Type", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, req)
	})
}
