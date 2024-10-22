package main

import (
	"log"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{
        Addr: "127.0.0.1:12345",
    }); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
