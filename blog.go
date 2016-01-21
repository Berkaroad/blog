package main

import (
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Router.Get("/", func(req *http.Request, w martini.ResponseWriter) (int, string) {
		return 200, "Hello Jerry!"
	})
	m.Run()
}
