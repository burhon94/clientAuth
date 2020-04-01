package main

import (
	"flag"
	"log"
	"net"
	"net/http"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9999", "Server port")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	serverUp(addr)
}

func serverUp(addr string) {
	log.Printf("Client Auth Service starting on: %s ...", addr)
	panic(http.ListenAndServe(addr, nil))
}
