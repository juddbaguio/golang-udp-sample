package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func handleWrite(conn net.PacketConn, payload []byte, addr net.Addr) {
	conn.WriteTo([]byte(fmt.Sprint("I received the following: ♥ ", string(payload))), addr)
}

func main() {
	conn, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	log.Println("starting UDP server")

	bufferSize := make([]byte, 1024)
	for {
		log.Println("waiting for requests")
		n, addr, err := conn.ReadFrom(bufferSize[:])
		if err != nil {
			log.Println(err)
			break
		}

		log.Println(addr)
		log.Println("You sent me this: ", string(bufferSize[:n]))

		// spawn go routine to enable concurrent responses when there are 2 or more clients
		go handleWrite(conn, bufferSize[:n], addr)
	}
}
