package server

func (s *Server) InitRoutes() {
	s.router.GET(
		"/api/status",
		s.handleHealth(),
	)

	s.router.POST(
		"/api/client/0",
		s.handleNewClient(),
	)

	s.router.POST(
		"/api/client/password",
		s.handleEditPass(),
		)
}
