package dst

import (
	//"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		addr := r.Header.Get("X-Real-IP")
		if addr == "" {
			addr = r.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = r.RemoteAddr
			}
		}
		if next != nil {
			next.ServeHTTP(w, r)
		}
		log.Printf("[%s] %s from %s in %v\n", r.Method, r.URL, addr, time.Since(start))
	}
	return http.HandlerFunc(fn)
}
