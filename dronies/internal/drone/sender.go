package drone

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

type Sender struct {
	Conn *nats.Conn
}

func (s *Sender) SendMessage(subject string, msg Message) error {
	b, err := encodeJson(msg)

	if err != nil {
		return WithLog(err)
	}
	log.Println(string(b))
	return s.Conn.Publish(subject, b)
}

func encodeJson(msg Message) ([]byte, error) {
	return json.MarshalIndent(msg, "", "\t")
}
