package server

import "github.com/burhon94/clientAuth/pkg/middleware/logger"

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
		"/api/token",
		s.handleToken(),
	)

	s.router.POST(
		"/api/client/password",
		s.handleEditPass(),
	)

	s.router.POST(
		"/api/client/avatar",
		s.handleEditAvatar(),
		logger.Logger("Client edit Avatar"),
	)

	s.router.POST(
		"/api/client/edit",
		s.handleEditClient(),
	)
}
