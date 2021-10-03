package logger

import (
	"net/http"
	"strings"
)

func (l logger) log() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		p := strings.Split(r.URL.Path, "/")
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte(p[1]))
	})
	http.ListenAndServe(":9090", nil)
}

type World struct {
	Message string `json:"message"`
}

type logger struct {
}
