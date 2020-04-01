package server

func (s *Server) InitRoutes() {
	s.router.GET(
		"/",
		s.handleIndexPage(),
	)
}
