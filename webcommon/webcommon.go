package webcommon

import (
	"context"
	"net/http"
	"net/url"
)

// Standard HTTP handler that we use, which should contain generic HTTP
// handler processing that we expect from all APIs
type StandardHandler[Q any, QP QueryPointer[Q]] func(
	ctx context.Context,
	query QP,
) (Response, *HandlerError)

// type MyStore[T Query] struct {
// 	data T
// }

// func (s *MyStore[T]) Handle(item T) {
// 	values := url.Values{}
// 	// u := &url.URL{
// 	// 	Scheme: "https",
// 	// 	Host:   "example.com",
// 	// 	Path:   "/path/to/endpoint",
// 	// }
// 	item.ParseQuery(values)
// }

// func Define() {
// 	var storeA = &MyStore[*QueryData]{}
// 	a := &QueryData{}

// 	storeA.Handle(a)
// }

// type Q struct {
// 	QueryData // Embedded QueryData type within Q
// }

func (h StandardHandler[Q, QP]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var query QP = new(Q)
	// var query Q
	// var q Q
	// var query Q
	// query := New[QueryData](QueryData{})
	if err := query.ParseQuery(r.URL.Query()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	h(r.Context(), query)

	w.WriteHeader(http.StatusOK)
}

type QueryPointer[QT any] interface {
	*QT
	Query
}

// Expected to be used to wrap handler function to be translated into `http.HandlerFunc`
// func Handle(h StandardHandler) http.HandlerFunc {
// 	return h.ServeHTTP
// }

// // Serves HTTP request using specified generic parameter
// func (h StandardHandler[Q]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// query := new(Q)
// 	// var query Q
// 	query := &QueryData{}
// 	if err := query.ParseQuery(r.URL.Query()); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}

// 	// var q Q
// 	res, err := h(r.Context(), query)
// 	if err != nil {
// 		switch err.Code {
// 		case http.StatusBadRequest:
// 			w.WriteHeader(http.StatusBadRequest)
// 		case http.StatusUnauthorized, http.StatusForbidden:
// 			w.WriteHeader(http.StatusForbidden)
// 		default:
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 		if err := json.NewEncoder(w).Encode(err); err != nil {
// 			log.Error().Stack().Err(err).Send()
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	if err := json.NewEncoder(w).Encode(res); err != nil {
// 		log.Error().Stack().Err(err).Send()
// 	}
// }

type QueryData struct {
	IDSS string
}

// var _ Query = &QueryData{}

func (q *QueryData) ParseQuery(values url.Values) error {
	q.IDSS = "everything ok"
	return nil
}

func New[T any](value T) *T {
	ptr := &value
	return ptr
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
