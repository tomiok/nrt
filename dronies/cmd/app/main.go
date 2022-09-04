package main

import (
	"dronies/cmd/internal/drone"
	"log"
	"time"
)

const maxDrones = 10

func main() {
	now := time.Now()
	run()
	log.Println(time.Since(now))
}

func run() error {
	drones := drone.DronesGeneration(maxDrones)

	ch := make(chan string)
	go drone.DataAggregator(ch, drones)

	for msg := range ch {
		log.Println(".." + msg)
	}

	return nil
}
