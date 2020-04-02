package server

import (
	"context"
	"errors"
	"github.com/burhon94/clientAuth/pkg/core/client"
	"github.com/burhon94/clientAuth/pkg/responses"
	readJSON "github.com/burhon94/json/cmd/reader"
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
		var bodyRequest client.NewClientStruct
		err := readJSON.ReadJSONHTTP(request, &bodyRequest)
		if err != nil {
			err := responses.SetResponseBadRequest(writer, "err.json_invalid")
			if err != nil {
				responses.InternalErr(writer)
			}
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Hour)
		err = s.clientSvc.NewClient(ctx, bodyRequest)
		if err != nil {
			log.Printf("can't create new client: %v", err)
			switch {
			case errors.Is(err, client.ErrBadRequest):
				err := responses.SetResponseBadRequest(writer, "err.bad_request")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			//case errors.Is(err, client.ErrLoginExist):
			//	err := responses.SetResponseBadRequest(writer, "err.login_exist")
			//	if err != nil {
			//		responses.InternalErr(writer)
			//	}
			//	return
			//
			//case errors.Is(err, client.ErrPhoneRegistered):
			//	err := responses.SetResponseBadRequest(writer, "err.phone_registered")
			//	if err != nil {
			//		responses.InternalErr(writer)
			//	}
			//	return

			default:
				err := responses.SetResponseInternalErr(writer, "err.internal_err")
				if err != nil {
					responses.InternalErr(writer)
				}
			}
		}
	}
}
