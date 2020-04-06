package server

import (
	"github.com/burhon94/clientAuth/pkg/core/client"
	"github.com/burhon94/clientAuth/pkg/middleware/authenticated"
	"github.com/burhon94/clientAuth/pkg/middleware/jwt"
	"github.com/burhon94/clientAuth/pkg/middleware/logger"
	"reflect"
)

func (s *Server) InitRoutes() {
	s.router.GET(
		"/api/status",
		s.handleHealth(),
		logger.Logger("Status Service"),
	)

	s.router.POST(
		"/api/client/0",
		s.handleNewClient(),
		authenticated.Authenticated(jwt.IsContextNonEmpty),
		jwt.JWT(reflect.TypeOf((*client.TokenPayload)(nil)).Elem(), s.secret),
		logger.Logger("New Client on Service"),
	)

	s.router.POST(
		"/api/token",
		s.handleToken(),
		logger.Logger("Generate Token"),
	)

	s.router.POST(
		"/api/client/password",
		s.handleEditPass(),
		authenticated.Authenticated(jwt.IsContextNonEmpty),
		jwt.JWT(reflect.TypeOf((*client.TokenPayload)(nil)).Elem(), s.secret),
		logger.Logger("Client edit Password"),
	)

	s.router.POST(
		"/api/client/avatar",
		s.handleEditAvatar(),
		authenticated.Authenticated(jwt.IsContextNonEmpty),
		jwt.JWT(reflect.TypeOf((*client.TokenPayload)(nil)).Elem(), s.secret),
		logger.Logger("Client edit Avatar"),
	)

	s.router.POST(
		"/api/client/edit",
		s.handleEditClient(),
		authenticated.Authenticated(jwt.IsContextNonEmpty),
		jwt.JWT(reflect.TypeOf((*client.TokenPayload)(nil)).Elem(), s.secret),
		logger.Logger("Client edit Data"),
	)
}
