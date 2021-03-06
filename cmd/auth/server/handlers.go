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

			case errors.Is(err, client.ErrTimeCtx):
				err := responses.SetResponseTimeOut(writer)
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			case errors.Is(err, client.ErrInternal):
				err := responses.SetResponseInternalErr(writer, "err.internal_err")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			default:
				err := responses.SetResponseBadRequest(writer, "err.unknown_err")
				if err != nil {
					responses.InternalErr(writer)
				}
				return
			}
		}

		err = responses.ResponseOK(writer)
		if err != nil {
			responses.InternalErr(writer)
		}
	}
}

func (s *Server) handleToken() http.HandlerFunc {
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

		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		token, err := s.clientSvc.GenerateToken(ctx, requestData)
		if err != nil {
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

			case errors.Is(err, client.ErrInternal):
				err := responses.SetResponseInternalErr(writer, "err.internal_error")
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

			case errors.Is(err, client.ErrTimeCtx):
				err := responses.SetResponseTimeOut(writer)
				if err != nil {
					responses.InternalErr(writer)
				}

			default:
				err := responses.SetResponseBadRequest(writer, "err.unknown_err")
				if err != nil {
					responses.InternalErr(writer)
				}
				return
			}
		}

		err = responses.SetResponseJSON(writer, &token)
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
		err = s.clientSvc.EditClientPass(ctx, request, bodyRequest)
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

			case errors.Is(err, client.ErrTimeCtx):
				err := responses.SetResponseTimeOut(writer)
				if err != nil {
					responses.InternalErr(writer)
				}

			case errors.Is(err, client.ErrInternal):
				err := responses.SetResponseInternalErr(writer, "err.internal_error")
				if err != nil {
					responses.InternalErr(writer)
				}

			default:
				err := responses.SetResponseBadRequest(writer, "err.unknown_err")
				if err != nil {
					responses.InternalErr(writer)
				}
				return
			}
		}

		err = responses.ResponseOK(writer)
		if err != nil {
			responses.InternalErr(writer)
		}
	}
}

func (s *Server) handleEditAvatar() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var dataRequest client.EditClientAvatar
		err := readJSON.ReadJSONHTTP(request, &dataRequest)
		if err != nil {
			err := responses.SetResponseBadRequest(writer, "err.json_invalid")
			if err != nil {
				responses.InternalErr(writer)
			}

			return
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		err = s.clientSvc.EditClientAvatar(ctx, request, dataRequest.AvatarUrl)
		if err != nil {
			switch {
			case errors.Is(err, client.ErrBadRequest):
				err := responses.SetResponseBadRequest(writer, "err.bad_request")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			case errors.Is(err, client.ErrInternal):
				err := responses.SetResponseInternalErr(writer, "err.internal_error")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			case errors.Is(err, client.ErrTimeCtx):
				err := responses.SetResponseTimeOut(writer)
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			default:
				err := responses.SetResponseBadRequest(writer, "err.unknown_error")
				if err != nil {
					responses.InternalErr(writer)
				}
				return
			}
		}

		err = responses.ResponseOK(writer)
		if err != nil {
			responses.InternalErr(writer)
		}
	}
}

func (s *Server) handleEditClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var clientRequest client.EditClient
		err := readJSON.ReadJSONHTTP(request, &clientRequest)
		if err != nil {
			err := responses.SetResponseBadRequest(writer, "err.invalid_json")
			if err != nil {
				responses.InternalErr(writer)
			}

			return
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		err = s.clientSvc.EditClient(ctx, request, clientRequest)
		if err != nil {
			switch {
			case errors.Is(err, client.ErrBadRequest):
				err := responses.SetResponseBadRequest(writer, "err.bad_request")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			case errors.Is(err, client.ErrInternal):
				err := responses.SetResponseInternalErr(writer, "err.internal_error")
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			case errors.Is(err, client.ErrTimeCtx):
				err := responses.SetResponseTimeOut(writer)
				if err != nil {
					responses.InternalErr(writer)
				}
				return

			default:
				err := responses.SetResponseBadRequest(writer, "err.unknown_error")
				if err != nil {
					responses.InternalErr(writer)
				}
				return
			}
		}

		err = responses.ResponseOK(writer)
		if err != nil {
			responses.InternalErr(writer)
		}
	}
}
