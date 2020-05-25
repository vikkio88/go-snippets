package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	host, port := getArgs()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Printf("%d CLOSED (%s)\n", port, err)
		return
	}
	conn.Close()
	log.Printf("%d OPEN\n", port)
}

func getArgs() (string, int) {
	host := flag.String("h", "", "-h host")
	port := flag.Int("p", 0, "-p port")
	flag.Parse()

	if *host == "" || *port <= 0 {
		panic("invalid input")
	}

	return *host, *port
}
