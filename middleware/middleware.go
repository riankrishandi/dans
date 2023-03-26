package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/riankrishandi/dans/auth"
	"github.com/riankrishandi/dans/render"
)

var (
	errAuthTokenRequired = errors.New("auth token is required")
)

type Middleware struct {
	renderer *render.Renderer
}

func New(r *render.Renderer) *Middleware {
	return &Middleware{
		renderer: r,
	}
}

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.TrimSpace(token)
		if token == "" {
			m.renderer.RenderJSON(w, http.StatusUnauthorized, errAuthTokenRequired)
			return
		}

		claims, err := auth.VerifyToken(token)
		if err != nil {
			m.renderer.RenderJSON(w, http.StatusForbidden, err)
			return
		}

		next.ServeHTTP(w, r.WithContext(auth.NewContext(r.Context(), claims)))
	})
}
