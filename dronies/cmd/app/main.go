package main

import (
	drone2 "dronies/internal/drone"
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

	sender := drone2.Sender{
		Conn: nc,
	}

	exec := drone2.Executor{
		Sender: &sender,
	}

	return &deps{
		Exec: &exec,
	}, nil
}

type deps struct {
	Exec *drone2.Executor
}
