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

func (s *Server) handleHealth() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("service ok!"))
		writer.WriteHeader(http.StatusOK)
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

			return
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Second)
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

			case errors.Is(err, client.ErrLoginExist):
				err := responses.SetResponseBadRequest(writer, "err.login_exist")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			case errors.Is(err, client.ErrPhoneRegistered):
				err := responses.SetResponseBadRequest(writer, "err.phone_registered")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			default:
				err := responses.SetResponseInternalErr(writer, "err.internal_err")
				if err != nil {
					responses.InternalErr(writer)
				}
			}
		}

		err = responses.ResponseOK(writer)
		if err != nil {
			responses.InternalErr(writer)
		}
	}
}

func (s *Server) handleSignIn() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var requestData client.SignIn
		err := readJSON.ReadJSONHTTP(request, &requestData)
		if err != nil {
			err := responses.SetResponseBadRequest(writer, "err.json_invalid")
			if err != nil {
				responses.InternalErr(writer)
			}

			return
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Hour)
		err = s.clientSvc.SignIn(ctx, requestData)
		switch {
		case errors.Is(err, client.ErrInvalidLogin):
			err := responses.SetResponseBadRequest(writer, "err.login_wrong")
			if err != nil {
				responses.InternalErr(writer)
			}
			return

		case errors.Is(err, client.ErrInvalidPassword):
			err := responses.SetResponseBadRequest(writer, "err.password_wrong")
			if err != nil {
				responses.InternalErr(writer)
			}
			return

		case errors.Is(err, client.ErrBadRequest):
			err := responses.SetResponseBadRequest(writer, "err.bad_request")
			if err != nil {
				responses.InternalErr(writer)
			}
			return

		}

		err = responses.ResponseOK(writer)
		if err != nil {
			responses.InternalErr(writer)
		}
	}
}

func (s *Server) handleEditPass() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var bodyRequest client.EditClientPass
		err := readJSON.ReadJSONHTTP(request, &bodyRequest)
		if err != nil {
			err := responses.SetResponseBadRequest(writer, "err.json_invalid")
			if err != nil {
				responses.InternalErr(writer)
			}

			return
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		err = s.clientSvc.EditClientPass(ctx, bodyRequest)
		if err != nil {
			switch {
			case errors.Is(err, client.ErrBadRequest):
				err := responses.SetResponseBadRequest(writer, "err.bad_request")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			case errors.Is(err, client.ErrInvalidPassword):
				err := responses.SetResponseBadRequest(writer, "err.password_wrong")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			default:
				err := responses.SetResponseInternalErr(writer, "err.internal_err")
				if err != nil {
					responses.InternalErr(writer)
				}
			}
		}

		err = responses.ResponseOK(writer)
		if err != nil {
			responses.InternalErr(writer)
		}
	}
}
