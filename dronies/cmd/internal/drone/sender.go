package drone

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

type Sender struct {
	conn *nats.Conn
}

func (s *Sender) SendMessage(subject string, msg interface{}) error {
	b, err := encode(msg)

	if err != nil {
		return WithLog(err)
	}

	return s.conn.Publish(subject, b)
}

func encode[T any](tt T) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(tt)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

type Message struct {
	DroneID uint64
	ScanAt  time.Time
}

func DataAggregator(ch chan string, drones []Drone) {
	for {
		for i := range drones {
			drones[i].ScanEnemy()
		}

		for _, d := range drones {
			for _, enemy := range d.Scans.Squad.Enemies {
				droneInfo := fmt.Sprintf("Drone %d position: %f  %f at: %s", d.ID, d.Lon, d.Lat, d.Scans.ScanTime.Format(time.ANSIC))
				enemyInfo := enemy.GetInfo()
				msg := fmt.Sprintf("%s - %s", droneInfo, enemyInfo)

				ch <- msg
			}
		}
		time.Sleep(1 * time.Second)
	}
}
