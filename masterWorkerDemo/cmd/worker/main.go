package main

import (
	"flag"
	"golangdemo/masterWorkerDemo/worker"
	"log"
	"os"
	"os/signal"
)

var configFile = flag.String("config", "../../conf/config.json", "app config file")

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	defer func() {
		if err := recover(); err != nil {
			log.Println("Worker recovered from err:", err)
			shutdown <- os.Interrupt
			log.Println("Worker was recovered from panic just now and is shutting down soon")
		}
	}()

	flag.Parse()
	conf, err := worker.LoadConfig(*configFile)
	if err != nil {
		log.Println("configur file load err: ", err)
		return
	}

	server, err := worker.GetTargetServer(conf)
	if err != nil {
		log.Println("Worker GetTargetServer err: ", err)
		return
	}

	if err = server.Start(); err != nil {
		log.Println("Worker server start failed")
		return
	}

	<-shutdown

	server.Stop()
}
