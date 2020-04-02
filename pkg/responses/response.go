package responses

import (
	writeJSON "github.com/burhon94/json/cmd/writer"
	"net/http"
	)

type ErrorDTO struct {
	Errors string `json:"errors"`
}

type Response struct {
	Respons string `json:"respons"`
}

func SetResponseBadRequest(writer http.ResponseWriter, errText string) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	err := writeJSON.WriteJSONHTTP(writer, &ErrorDTO{
		errText,
	})

	return err
}

func SetResponseInternalErr(writer http.ResponseWriter, errText string) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	err := writeJSON.WriteJSONHTTP(writer, &ErrorDTO{
		errText,
	})

	return err
}

func BadRequest(writer http.ResponseWriter) {
	http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func InternalErr(writer http.ResponseWriter)  {
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ResponseOK(writer http.ResponseWriter) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := writeJSON.WriteJSONHTTP(writer, &Response{
		"OK",
	})

	return err
}