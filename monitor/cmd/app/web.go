package main

import (
	"fmt"
	"log"
	"monitor/internal/monitor"
	"net/http"
)

type SSE struct {
	MessageCh chan monitor.Message
}

func (s *SSE) EventHandler(w http.ResponseWriter, _ *http.Request) {
	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "streaming unsupported!", http.StatusInternalServerError)
		return
	}
	setHeaders(w)

	for {
		_, err := fmt.Fprintf(w, "%s \n", func(m monitor.Message) string {
			droneID := m.DroneID
			enemies := len(m.Enemies)
			position := fmt.Sprintf("%.2f | %.2f ", m.Position.Lon, m.Position.Lat)
			scannedAt := m.ScanAt

			return fmt.Sprintf("droneID: %d, enemy count:%d, position:%s, time:%s",
				droneID,
				enemies,
				position,
				scannedAt,
			)
		}(<-s.MessageCh))
		if err != nil {
			log.Println(err)
		}

		flusher.Flush()
	}
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Transfer-Encoding", "chunked")

}
