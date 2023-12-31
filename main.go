// Based on https://github.com/zooraze/chi-vue-spa
// and https://github.com/thefrol/go-chi-yandex-cloud-template

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var Router = chi.NewRouter()

func init() {
	fileServer := http.FileServer(http.Dir("web"))
	Router.Handle("/*", fileServer)
}

func main() {
	http.ListenAndServe(":8080", Router)
}
