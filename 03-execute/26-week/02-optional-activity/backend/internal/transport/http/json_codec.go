package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"workshop/internal/domain"
)

// maxRequestByte bounds the body the API accepts, so a large upload cannot
// exhaust memory before the handler even looks at it.
const maxRequestByte = 1 << 20

// decode reads a JSON body into the target. A malformed body is reported as a
// domain input error, never as a decoder message that leaks internals.
func decode(writer http.ResponseWriter, request *http.Request, target any) error {
	request.Body = http.MaxBytesReader(writer, request.Body, maxRequestByte)
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("%w: el cuerpo de la peticion no es valido", domain.ErrInvalidInput)
	}
	return nil
}

// respond writes a JSON payload with the given status code.
func respond(writer http.ResponseWriter, status int, payload any) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(status)
	if payload == nil {
		return
	}
	_ = json.NewEncoder(writer).Encode(payload)
}
