package drone

import (
	"log"
	"time"
)

const maxDrones = 10

type Executor struct {
	Sender *Sender
}

func (e *Executor) Run() {
	drones := DronesGeneration(maxDrones)

	ch := make(chan Message)

	go Execute(ch, drones)

	var i int
	for msg := range ch {
		err := e.Sender.SendMessage("enemies", msg)
		if err != nil {
			log.Println(err)
		}
		log.Println("message sent ", i)
		i++
	}
}

func Execute(ch chan Message, drones []Drone) {
	for {
		ticker := time.NewTicker(time.Duration(random.Intn(10)+1) * time.Second)
		select {
		case <-ticker.C:
			for i := range drones {
				drone := drones[i]
				drone.ScanEnemy()
				ch <- drone.Read()
			}
		}
	}
}
