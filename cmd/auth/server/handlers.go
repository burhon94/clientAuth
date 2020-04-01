package server

import (
	"log"
	"net/http"
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
