package main

import (
	"flag"
	"fmt"
	"github.com/burhon94/alifMux/pkg/mux"
	"github.com/burhon94/bdi/pkg/di"
	"github.com/burhon94/clientAuth/cmd/auth/server"
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
	container := di.NewContainer()

	err := container.Provide(
		server.NewServer,
		mux.NewExactMux,
		)
	if err != nil {
		panic(fmt.Sprintf("can't set provide: %v", err))
	}

	container.Start()
	var appServer *server.Server
	container.Component(&appServer)

	log.Printf("Client Auth Service starting on: %s ...", addr)
	panic(http.ListenAndServe(addr, appServer))
}
