package flagshelper

import "flag"

func GetArgs() (string, int) {
	host := flag.String("h", "", "-h host")
	port := flag.Int("p", 0, "-p port")
	flag.Parse()

	if *host == "" || *port <= 0 {
		panic("invalid input")
	}

	return *host, *port
}
