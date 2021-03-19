package goget

import (
	"net/http"
	"strings"
)

type Package struct {
	VCS string
	URL string
}

type View struct {
	Canonical  string
	Repository string
	Request    string
	VCS        string
}

type Packages map[string]*Package

func (m Packages) Get(r *http.Request) *View {
	url := r.Host + r.URL.Path

	for k, p := range m {
		s := strings.TrimPrefix(url, k)

		if s == "" || s[0] == '/' {
			v := &View{
				Request:    url,
				Canonical:  k,
				Repository: p.URL,
				VCS:        p.VCS,
			}

			return v
		}
	}

	return nil
}

func (m Packages) SetDefaults() error {
	for k, v := range m {
		if v.URL == "" {
			delete(m, k)
			continue
		}

		if v.VCS == "" {
			v.VCS = "git"
		}
	}

	return nil
}
