package webcommon

import (
	"context"
	"net/http"
	"net/url"
)

// Standard HTTP handler that we use, which should contain generic HTTP
// handler processing that we expect from all APIs
type StandardHandler[Q any, QP QueryPointer[Q], R Response] func(
	ctx context.Context,
	query QP,
) (R, *HandlerError)

func Handle[
	Q any,
	QP QueryPointer[Q],
	R Response,
](h StandardHandler[Q, QP, R]) http.HandlerFunc {
	return h.ServeHTTP
}

func (h StandardHandler[Q, QP, R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var query QP = new(Q)
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
