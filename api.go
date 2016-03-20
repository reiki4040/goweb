package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zenazn/goji/web"
)

type API struct{}

func (api *API) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I'm goji!\n")
}

func (api *API) Sleep(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "sleeping...")
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "sleeped 5 sec")
}

func MyMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logging my middleware start.")
		h.ServeHTTP(w, r)
		fmt.Println("logging my middleware end.")
	}

	return http.HandlerFunc(fn)
}
