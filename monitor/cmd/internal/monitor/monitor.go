package monitor

import (
	"github.com/nats-io/nats.go"
	"log"
)

type Monitor struct {
	Conn *nats.Conn
}

func (m *Monitor) Read() error {
	s, err := m.Conn.Subscribe("enemies_1", func(msg *nats.Msg) {
		log.Println(string(msg.Data))
	})

	if err != nil {
		return err
	}

	return s.Drain()
}
