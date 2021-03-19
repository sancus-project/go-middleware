package goget

import (
	"log"
	"net/http"
)

func MiddlewareFromFile(fn string, renderer Renderer) func(http.Handler) http.Handler {
	m, err := PackagesFromFile(fn)

	if err != nil {
		log.Fatal(err)
	}

	return m.Middleware(renderer)
}

func (m *Packages) Middleware(renderer Renderer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return m.Handler(next, renderer)
	}
}

func (m *Packages) Handler(next http.Handler, renderer Renderer) http.Handler {
	if renderer == nil {
		renderer = DefaultRenderer()
	}

	if next == nil {
		next = http.NotFoundHandler()
	}

	fn := func(w http.ResponseWriter, r *http.Request) {

		pkg := m.Get(r)
		if pkg != nil {
			err := renderer(w, r, pkg)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
