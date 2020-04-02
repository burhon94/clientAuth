package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

func (s *Server) handleIndexPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Server is up success!"))
		if err != nil {
			log.Printf("can't write: %v", err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (s *Server) handleNewClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(request.Context(), time.Hour)
		err := s.clientSvc.NewClient(ctx, "testName", "testLastName", "testLogin", "testPass", request)
		if err != nil {
			log.Printf("can't create new client: %v", err)
		}
	}
}
