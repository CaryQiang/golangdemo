package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "172.20.12.225:8001")
	if err != nil {
		log.Fatalf("client: dial: %s\n", err)
	}

	defer conn.Close()

	log.Println("client: connected to: ", conn.RemoteAddr())

	message := "Hello, im clent \n"
	n, err := conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("clent: write: %s", err)
	}

	reply := make([]byte, 512)
	n, err = conn.Read(reply)
	if err != nil {
		log.Fatalf("clent: read: %s", err)
	}

	log.Println("clent: read: ", string(reply[:n]))

}
