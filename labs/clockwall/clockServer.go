// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func handleConn(c net.Conn, timezone string) {
	defer c.Close()
	for {
		loc, _ := time.LoadLocation(timezone)
		_, err := io.WriteString(c, timezone+":\t"+time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	timezone := os.Getenv("TZ")
	port := flag.Int("port", 9000, "port number.")
	flag.Parse()

	fmt.Printf("Timezone: %v\n", timezone)
	fmt.Printf("Port: %v\n", *port)

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, timezone) // handle connections concurrently
	}
}
