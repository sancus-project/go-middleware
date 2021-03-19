package goget

import (
	"html/template"
	"log"
	"net/http"
)

type Renderer func(w http.ResponseWriter, r *http.Request, v *View) error

func DefaultRenderer() Renderer {

	tmpl := `<!DOCTYPE html>
<head>
	<meta name="go-import" content="{{.Canonical}} {{.VCS}} {{.Repository}}" />
	<meta http-equiv="refresh" content="5; url=https://pkg.go.dev/{{.Request}}" />
</head>
<body>
	<pre>git clone <a href="{{.Repository}}">{{.Repository}}</a></pre>
	<pre>go get <a href="https://pkg.go.dev/{{.Request}}">{{.Request}}</a></pre>
	<pre>import "<a href="https://pkg.go.dev/{{.Request}}">{{.Request}}</a></pre>
</body>
`
	t, err := template.New("package").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	fn := func(w http.ResponseWriter, _ *http.Request, v *View) error {
		return t.Execute(w, v)
	}

	return fn
}
