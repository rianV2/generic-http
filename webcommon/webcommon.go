package webcommon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
)

// Expected to be used to wrap handler function to be translated into `http.HandlerFunc`
func Handle[
	Q Query,
	U QueryPointer[Q],
	R Response,
](h StandardHandler[Q, U, R]) http.HandlerFunc {
	return h.ServeHTTP
}

type CustomHandler struct {
	// You can add any fields or configurations specific to your handler here
}

// Implementing the ServeHTTP method for CustomHandler
func (ch *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Your logic for handling the HTTP request goes here
	fmt.Fprintf(w, "Hello, this is a custom handler!")
}

type QueryPointer[QT any] interface {
	*QT
	Query
}

// Standard HTTP handler that we use, which should contain generic HTTP
// handler processing that we expect from all APIs
type StandardHandler[
	Q Query,
	U QueryPointer[Q],
	R Response,
] func(
	ctx context.Context,
	query U,
) (R, *HandlerError)

// Serves HTTP request using specified generic parameter
func (h StandardHandler[Q, U, R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := U(new(Q))
	if err := query.ParseQuery(r.URL.Query()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// var q Q
	res, err := h(r.Context(), query)
	if err != nil {
		switch err.Code {
		case http.StatusBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		case http.StatusUnauthorized, http.StatusForbidden:
			w.WriteHeader(http.StatusForbidden)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Error().Stack().Err(err).Send()
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Error().Stack().Err(err).Send()
	}
}

type QueryData struct {
	IDSS string
}

// var _ Query = &QueryData{}

func (q *QueryData) ParseQuery(values url.Values) error {
	q.IDSS = "everything ok"
	return nil
}

// =====================================================
type Response interface {
	ToResponse(http.ResponseWriter) ([]byte, error)
}

type ResponseEmpty struct{}

var _ Response = &ResponseEmpty{}

func (ResponseEmpty) ToResponse(http.ResponseWriter) ([]byte, error) {
	return nil, nil
}

type HandlerError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (h *HandlerError) Error() string {
	return h.Message
}

type Query interface {
	ParseQuery(values url.Values) error
}

type QueryEmpty struct{}

var _ Query = QueryEmpty{}

func (QueryEmpty) ParseQuery(url.Values) error {
	return nil
}
