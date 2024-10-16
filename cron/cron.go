package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := "*/1 * * * * *"
	c.AddFunc(spec, func() {
		i ++
		log.Println("cron1 running:", i)
	})
	c.AddFunc(spec, func() {
		i ++
		log.Println("cron2 running:", i)
	})
	c.Start()
	defer c.Stop()
	select {}
}