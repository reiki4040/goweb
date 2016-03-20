package main

import (
	"flag"
	"net"
	"net/http"
	"os"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var (
	optFd uint
)

func init() {
	// file descriptor option for Circus
	flag.UintVar(&optFd, "fd", 0, "File descriptor to listen and serve.")

	// hiding goji -bind option.
	flag.Parse()
}

func main() {
	api, err := createAPI()
	if err != nil {
		panic(err)
	}

	rootMux := goji.DefaultMux
	rootMux.Handle("/api/*", api)

	if optFd != 0 {
		l, err := net.FileListener(os.NewFile(uintptr(optFd), ""))
		if err != nil {
			panic(err)
		}

		goji.ServeListener(l)
	} else {
		// if not specified fd, then goji default(:8000 or -bind arg)
		goji.Serve()
	}
}

func createAPI() (http.Handler, error) {
	api := &API{}

	apiMux := web.New()
	apiMux.Handle("/api/hello", api.Hello)
	apiMux.Handle("/api/sleep", api.Sleep)
	apiMux.Use(MyMiddleware)

	return apiMux, nil
}
