# Simple near real time project
This is a simple demo of 2 servers communicating in near real time

### Server Side Events (sse) in Go
Using Golang and nats to provide SSE

### Run local

run nats
```bash
docker-compose up
```

run dronies
```bash
go run dronies/cmd/app/*.go
```

run monitor
```bash
go run monitor/cmd/app/*.go
```

### test in browser
go to: localhost:3335/listen