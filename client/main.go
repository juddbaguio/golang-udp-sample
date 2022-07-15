package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("udp", ":8080")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	buffer := make([]byte, 1024)
	for i := 0; i < 1e9; i++ {
		_, err := conn.Write([]byte(fmt.Sprint(i)))

		if err != nil {
			log.Println(err)
			continue
		}

		n, err := conn.Read(buffer)
		log.Println(string(buffer[:n]))
		if err != nil {
			log.Println(err)
			continue
		}

		time.Sleep(time.Second)

	}
}
