package main

import (
	"fmt"
	"log"
	"net/http"
)

type SSE struct {
	WebCh chan []byte
}

func (s *SSE) EventHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	setHeaders(w)

	for {
		_, err := fmt.Fprintf(w, "hello %s", <-s.WebCh)
		if err != nil {
			log.Println(err)
		}

		flusher.Flush()
	}
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	//w.Header().Set("Content-Type", "application/x-ndjson")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Transfer-Encoding", "chunked")
}
