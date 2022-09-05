package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"monitor/cmd/internal/monitor"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type dependenciess struct {
	monitor monitor.Monitor
}

func deps() {

}

func run() error {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		return err
	}

	m := monitor.Monitor{
		Conn: nc,
	}

	for {
		err := m.Read()

		if err != nil {
			break
		}
	}

	return nil
}
