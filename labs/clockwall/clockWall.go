package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func getTime(address string, ch chan string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return // e.g., wrong address
	}
	defer conn.Close()

	buffer := make([]byte, 1024) // only a string with timezone + time. 1024 SHOULD be enough
	for {
		n, err := conn.Read(buffer)
		if n > 0 { // n, number of bytes readed
			message := string(buffer[:n])
			ch <- message
		}

		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	ch := make(chan string, len(os.Args[1:]))

	for _, argument := range os.Args[1:] {
		address := strings.Split(argument, "=")[1] // follows the format: TIMEZONE=ADDRESS
		go getTime(address, ch)
	}

	for time := range ch {
		fmt.Printf("%v", time)
	}
}
