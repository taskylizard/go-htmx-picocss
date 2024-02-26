package main

import (
	"github.com/donseba/go-htmx"
	"github.com/donseba/go-htmx/middleware"
	"log"
	"net/http"
)

type (
	App struct {
		htmx *htmx.HTMX
	}

	route struct {
		path    string
		handler http.Handler
	}
)

func main() {
	// new app with htmx instance
	app := &App{
		htmx: htmx.New(),
	}

	mux := http.NewServeMux()

	// wrap routes with the middleware
	wrapRoutes(mux, middleware.MiddleWare, []route{
		{path: "/", handler: http.HandlerFunc(app.Home)},
		{path: "/pico.css", handler: http.HandlerFunc(app.PicoCSS)},
	})

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (a *App) PicoCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pico.css")
}

// wrapRoutes takes a mux, middleware, and a list of routes to apply the middleware to.
func wrapRoutes(mux *http.ServeMux, mw func(http.Handler) http.Handler, routes []route) {
	for _, r := range routes {
		mux.Handle(r.path, mw(r.handler))
	}
}
