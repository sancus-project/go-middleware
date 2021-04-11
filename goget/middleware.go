package goget

import (
	"log"
	"net/http"
)

// Middleware view
type Middleware struct {
	Renderer          Renderer // renders go-import html
	OnlyGoGet         bool     // only render go-import html if ?go-get=1
	RedirectToDoc     bool     // if package exists, but ?go-get=1 isn't give, redirect to documentation
	RedirectToSources bool     // if package exists, but ?go-get=1 isn't give, redirect to sources
}

// Middleware constructor from .ini file
func MiddlewareFromFile(fn string, options ...MiddlewareOption) func(http.Handler) http.Handler {
	m, err := PackagesFromFile(fn)

	if err != nil {
		log.Fatal(err)
	}

	return m.NewMiddleware(options...)
}

// Middleware constructor from Config
func (c *Config) NewMiddleware(options ...MiddlewareOption) func(http.Handler) http.Handler {
	v := NewMiddleware(c.Renderer, options...)
	fn := func(next http.Handler) http.Handler {
		return v.NewHandler(next, c.config.Package)
	}

	return fn
}

// Middleware constructor from Packages
func (m Packages) NewMiddleware(options ...MiddlewareOption) func(http.Handler) http.Handler {
	v := NewMiddleware(nil, options...)
	fn := func(next http.Handler) http.Handler {
		return v.NewHandler(next, m)
	}

	return fn
}

// Handler
func (v *Middleware) NewHandler(next http.Handler, m Packages) http.Handler {

	if next == nil {
		next = http.NotFoundHandler()
	}

	fn := func(w http.ResponseWriter, r *http.Request) {

		if pkg, goget := v.Get(m, r); pkg != nil {

			if goget {
				// render go-import page
				err := v.Renderer(w, r, pkg)
				if err != nil {
					log.Fatal(err)
				}
				return
			} else if v.RedirectToDoc {
				// redirect to doc
				url := "https://pkg.go.dev/" + pkg.Request
				http.Redirect(w, r, url, 302)
				return
			} else if v.RedirectToSources {
				// redirect to repo
				url := pkg.Repository
				http.Redirect(w, r, url, 302)
				return
			}
		}

		// next
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (_ Middleware) isGoGet(query map[string][]string) bool {
	// exactly ?go-get=1
	if len(query) == 1 {
		if m, ok := query["go-get"]; ok {
			if len(m) == 1 && m[0] == "1" {
				return true
			}
		}
	}
	return false
}

func (v *Middleware) Get(m Packages, r *http.Request) (pkg *View, goget bool) {

	pkg = m.Get(r)
	if pkg == nil {
		// unknown
	} else if !v.OnlyGoGet {
		// render go-import page to all matches
		goget = true
	} else {
		// render go-import page only if ?go-get=1 is passed. exactly.
		goget = v.isGoGet(r.URL.Query())
	}
	return
}
