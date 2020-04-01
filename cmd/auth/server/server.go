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
	//_, err := s.pool.Exec(context.Background(), dl.ClientDDL)
	//if err != nil {
	//	panic(fmt.Sprintf("can't init DB: %v", err))
	//}
	//
	//_, err = s.pool.Exec(context.Background(), dl.ClientDML)
	//if err != nil {
	//	panic(fmt.Sprintf("can't set DB: %v", err))
	//}

	s.InitRoutes()
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}
