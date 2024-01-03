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
	R Response,
](h StandardHandler[Q, R]) http.HandlerFunc {
	return h.ServeHTTP
}

type QueryPointer[QT any] interface {
	*QT
	Query
}

// Standard HTTP handler that we use, which should contain generic HTTP
// handler processing that we expect from all APIs
type StandardHandler[
	Q Query,
	R Response,
] func(
	ctx context.Context,
	query Q,
) (R, *HandlerError)

// Serves HTTP request using specified generic parameter
func (h StandardHandler[Q, R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// t := reflect.TypeOf((*Q)(nil))
	// Get the type of Q
	// queryType := reflect.TypeOf((*Q)(nil)).Elem()

	// Create a new instance of the query type using reflection
	// query := reflect.New(queryType).Interface()

	// q, _ := query.(QueryData)

	// query := reflect.New(t).Interface()

	// var qPointer QT = new(Q)
	var query Q
	if err := query.ParseQuery(r.URL.Query()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// // Check if the created instance implements the Query interface
	// if q, ok := query.(QueryData); ok {
	// 	// Call the ParseQuery method on the query instance
	// 	if err := q.ParseQuery(r.URL.Query()); err != nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		w.Write([]byte(err.Error()))
	// 		return
	// 	}
	// } else {
	// 	// Handle case where the created instance does not implement Query interface
	// 	// This might occur if the provided Q type does not adhere to the Query interface
	// 	// qType := reflect.TypeOf(query)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	// w.Write([]byte(fmt.Sprintf("Invalid query type %T", query)))
	// 	return
	// }

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

var _ Query = &QueryData{}

func (q *QueryData) ParseQuery(values url.Values) error {
	// fmt.Println("called")
	// q.IDSS = "everything ok"
	fmt.Println(q, "pasti nil")
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
