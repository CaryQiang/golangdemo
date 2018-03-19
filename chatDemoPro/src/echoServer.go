package main

import (
	"io"
	"log"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		log.Print("server: conn: waiting")
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("server: conn: read :%s\n", err)
			}
			break
		}

		log.Printf("server: conn: echo %q\n", string(buf[:n]))
		n, err = conn.Write([]byte("responds...end"))
		if err != nil {
			log.Printf("server: write: %s", err)
			break
		}

	}

	log.Println("server: conn: closed")
}

func main() {
	/*cert, err := tls.LoadX509KeyPair("/home/sensetime/Documents/rui.crt", "/home/sensetime/Documents/rui.key")
	if err != nil {
		log.Fatalf("server: loadKeys: %s", err)
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = time.Now
	config.Rand = rand.Reader
	service := ":8001"*/
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("server: listen: %s\n", err)
	}
	/*listener, err := conn.Listen("tcp", service, &config)
	if err != nil {
		log.Fatal("server: listen: %s", err)
	}
	*/
	log.Print("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}

		log.Printf("server: accepted from %s", conn.RemoteAddr())

		go handleClient(conn)

	}
}
