package logger

import (
	"net/http"
)

var message string

func (l Logger) Log(logMessage string) {
	if l.enabled {
		message = message + "\n" + string(logMessage)
	}
}

func (l Logger) Show() {
	if l.enabled {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "text/plain")
			rw.Write([]byte(message))
		})
		http.ListenAndServe(":"+l.port, mux)
	}
}

type Logger struct {
	port    string
	enabled bool
}
