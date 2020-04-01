package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/burhon94/alifMux/pkg/mux"
	"github.com/burhon94/bdi/pkg/di"
	"github.com/burhon94/clientAuth/cmd/auth/server"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"net/http"
)

type DSN string

// -host 0.0.0.0 -port 9999 -dsn postgres://user:pass@localhost:5555/client-auth
var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9999", "Server port")
	dsn  = flag.String("dsn", "postgres://user:pass@localhost:5555/client-auth", "Server DSN")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	serverUp(addr, *dsn)
}

func serverUp(addr string, dsn string) {
	container := di.NewContainer()

	err := container.Provide(
		server.NewServer,
		mux.NewExactMux,
		func() DSN { return DSN(dsn) },
		func(dsn DSN) *pgxpool.Pool {
			pool, err := pgxpool.Connect(context.Background(), string(dsn))
			if err != nil {
				panic(fmt.Errorf("can't create pool: %w", err))
			}

			return pool
		},
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
