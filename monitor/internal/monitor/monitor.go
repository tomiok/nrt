package monitor

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

type Monitor struct {
	Conn   *nats.Conn
	Broker chan []byte
}

func NewMonitor(conn *nats.Conn) *Monitor {
	m := &Monitor{
		Conn:   conn,
		Broker: make(chan []byte),
	}

	return m
}

func (m *Monitor) Listen(webCh chan Message) {
	for {
		select {
		case msg := <-m.Broker:
			var res Message
			err := json.Unmarshal(msg, &res)
			if err != nil {
				log.Println(err)
			}

			webCh <- res
		}
	}
}

func (m *Monitor) ReadForever() {
	for {
		err := m.Read()

		if err != nil {
			break
		}
	}
}

func (m *Monitor) Read() error {
	s, err := m.Conn.Subscribe("enemies", func(msg *nats.Msg) {
		log.Println("enemies report received")
		m.Broker <- msg.Data
	})

	if err != nil {
		return err
	}

	_ = s.Drain()

	return nil
}
