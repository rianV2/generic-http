package main

import (
	"context"
	"fmt"
	"net/http"

	wc "generic-http/webcommon"

	"github.com/go-chi/chi/v5"
)

// func main() {
// 	main1()
// }

func main() {

	// Create a new Router instance from chi
	r := chi.NewRouter()

	// Register the handler function to a specific route pattern using chi
	// customHandler := &wc.CustomHandler{}
	// custom := wc.StandardHandler[*wc.QueryData, wc.QueryPointer[*wc.QueryData]](GetCategory[*wc.QueryData]).ServeHTTP
	custom := wc.StandardHandler[wc.QueryData, *wc.QueryData](GetCategory).ServeHTTP

	r.Get("/categories/{id}", custom)

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

// func GetCategory2[
// 	Q wc.Query,
// 	U wc.QueryPointer[Q],
// 	R wc.Response,
// ](
// 	ctx context.Context,
// 	q U,
// ) (wc.ResponseEmpty, *wc.HandlerError) {
// 	// queryData, ok := q.(*wc.QueryData)
// 	// if !ok {
// 	// 	return wc.ResponseEmpty{}, &wc.HandlerError{
// 	// 		Code:    http.StatusBadRequest,
// 	// 		Message: "Invalid query type",
// 	// 	}
// 	// }

// 	// fmt.Println("getCategory", queryData)
// 	return wc.ResponseEmpty{}, nil
// }

func GetCategory(
	ctx context.Context,
	q *wc.QueryData,
) (wc.Response, *wc.HandlerError) {
	// queryData, ok := q.(*wc.QueryData)
	// if !ok {
	// 	return wc.ResponseEmpty{}, &wc.HandlerError{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "Invalid query type",
	// 	}
	// }
	fmt.Println(q)

	fmt.Println("getCategory")
	return wc.ResponseEmpty{}, nil
}
