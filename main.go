package main

import (
	"context"
	"fmt"
	"net/http"

	wc "generic-http/webcommon"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Create a new Router instance from chi
	r := chi.NewRouter()

	// Register the handler function to a specific route pattern using chi
	r.Get("/categories/{id}", wc.Handle(GetCategory))

	// Start the HTTP server using chi's Router
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

// func GetCategory(
// 	ctx context.Context,
// 	v wc.QueryData,
// 	q *wc.QueryData,
// ) (wc.ResponseEmpty, *wc.HandlerError) {
// 	// idssValue := q
// 	fmt.Println(q, "what")
// 	return wc.ResponseEmpty{}, nil
// }

func GetCategory(
	ctx context.Context,
	q *wc.QueryData,
) (wc.ResponseEmpty, *wc.HandlerError) {
	// data, err := usecase.ListCategory(ctx)
	// if err != nil {
	// 	return nil, &wc.HandlerError{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: err.Error(),
	// 	}
	// }

	// fmt.Println(q.IDSS)
	return wc.ResponseEmpty{}, nil
}
