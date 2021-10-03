package logger

import (
	"net/http"
	"strings"
)

func (l Logger) Log() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		p := strings.Split(r.URL.Path, "/")
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte(p[1]))
	})

	http.ListenAndServe(":8080", nil)
}

type Logger struct {
}
