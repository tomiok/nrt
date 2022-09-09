package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"monitor/cmd/internal/monitor"
	"net/http"
)

func main() {
	d, err := deps()
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr: ":3335",
	}

	http.HandleFunc("/listen", d.sse.EventHandler)
	go d.m.ReadForever()

	log.Fatal(server.ListenAndServe())
}

type dependencies struct {
	sse SSE
	m   *monitor.Monitor
}

func deps() (*dependencies, error) {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		return nil, err
	}

	ch := make(chan monitor.Message)
	m := monitor.NewMonitor(nc)

	go m.Listen(ch)
	sse := SSE{
		MessageCh: ch,
	}

	return &dependencies{
		sse: sse,
		m:   m,
	}, nil
}
