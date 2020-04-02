package server

func (s *Server) InitRoutes() {
	s.router.GET(
		"/",
		s.handleIndexPage(),
	)

	s.router.POST(
		"/api/client/0",
		s.handleNewClient(),
		)
}
