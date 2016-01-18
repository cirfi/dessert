package dst

import (
	"log"
	"net/http"
)

func Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Panic] %+v", err)
				w.Header().Set("Status", "500")
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		if next != nil {
			next.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(fn)
}
