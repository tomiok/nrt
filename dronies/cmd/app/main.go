package main

import (
	"dronies/cmd/internal/drone"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	now := time.Now()

	deps, err := buildDeps()

	if err != nil {
		log.Fatal(err)
	}

	deps.Exec.Run()

	log.Println(time.Since(now))
}

func buildDeps() (*deps, error) {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		return nil, err
	}

	sender := drone.Sender{
		Conn: nc,
	}

	exec := drone.Executor{
		Sender: &sender,
	}

	return &deps{
		Exec: &exec,
	}, nil
}

type deps struct {
	Exec *drone.Executor
}
