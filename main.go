package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

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
	r.Get("/categories/{id}", wc.Handle[GetCategoryQ, *GetCategoryQ](GetCategory))

	// Start the HTTP server using chi's Router
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

var _ wc.Query = &GetCategoryQ{}

type GetCategoryQ struct {
	IDSS string
}

func (q *GetCategoryQ) ParseQuery(values url.Values) error {
	q.IDSS = values.Get("idss")
	return nil
}
func GetCategory(
	ctx context.Context,
	q *GetCategoryQ,
) (wc.Response, *wc.HandlerError) {
	fmt.Println("GetCategory")
	fmt.Println(q.IDSS, "|retrieved data")

	return wc.ResponseEmpty{}, nil
}
