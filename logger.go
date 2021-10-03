package logger

import (
	"net/http"
)

var message string

func (l Logger) Log(logMessage string) {
	message = message + "\n" + string(logMessage)
}

func (l Logger) Show() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/plain")
		rw.Write([]byte(message))
	})
	http.ListenAndServe(":8080", mux)
}

type Logger struct {
}
