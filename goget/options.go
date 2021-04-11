package goget

// Middleware Options
type MiddlewareOption interface {
	isMiddlewareOption() MiddlewareOption
}

type middlewareOption struct{}

func (o middlewareOption) isMiddlewareOption() MiddlewareOption {
	return o
}

type OnlyGoGet struct {
	middlewareOption
}

type SetRenderer struct {
	middlewareOption
	Renderer Renderer
}

type RedirectToDoc struct {
	middlewareOption
}

type RedirectToSources struct {
	middlewareOption
}

func NewMiddleware(renderer Renderer, options ...MiddlewareOption) Middleware {

	v := Middleware{}

	for _, o := range options {
		if _, ok := o.(OnlyGoGet); ok {
			v.OnlyGoGet = true
		}
		if _, ok := o.(RedirectToDoc); ok {
			v.RedirectToDoc = true
			v.RedirectToSources = false
		}
		if _, ok := o.(RedirectToSources); ok {
			v.RedirectToDoc = false
			v.RedirectToSources = true
		}
		if r, ok := o.(SetRenderer); ok {
			v.Renderer = r.Renderer
		}
	}

	if v.Renderer == nil {
		if renderer == nil {
			renderer = DefaultRenderer()
		}

		v.Renderer = renderer
	}

	return v
}
