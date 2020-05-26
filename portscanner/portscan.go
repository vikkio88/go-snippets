package main

import (
	"fmt"
	"log"
	"net"

	fh "github.com/vikkio88/go-snippets/flagshelper"
)

func main() {
	host, port := fh.GetArgs()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Printf("%d CLOSED (%s)\n", port, err)
		return
	}
	conn.Close()
	log.Printf("%d OPEN\n", port)
}
