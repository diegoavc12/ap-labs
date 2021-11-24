// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
	"os"
)

func handleConn(c net.Conn) {
	defer c.Close()
	/*_, err := io.WriteString(c, os.Getenv("TZ")+": "+time.Now().Format("15:04:05\n"))
	if err != nil {
		return // e.g., client disconnected
	}
	time.Sleep(1 * time.Second)*/

	for{
		_, err := io.WriteString(c, os.Getenv("TZ")+": "+time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
	
}

func main() {
	port := flag.String("port", "9090", "number of tcp port")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+ *port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}
