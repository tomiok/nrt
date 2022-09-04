package drone

import (
	"bytes"
	"encoding/gob"
	"github.com/nats-io/nats.go"
)

type Sender struct {
	Conn *nats.Conn
}

func (s *Sender) SendMessage(subject string, msg Message) error {
	b, err := encode(msg)

	if err != nil {
		return WithLog(err)
	}

	return s.Conn.Publish(subject, b)
}

func encode[T any](tt T) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(tt)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
