package server

import (
	"github.com/burhon94/alifMux/pkg/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type Server struct {
	router *mux.ExactMux
	pool   *pgxpool.Pool
}

func NewServer(router *mux.ExactMux, pool *pgxpool.Pool) *Server {
	return &Server{router: router, pool: pool}
}

func (s *Server) Start() {
	s.InitRoutes()
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}
