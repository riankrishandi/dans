package render

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Renderer struct{}

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func New() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var res Response

	if err, ok := data.(error); ok {
		res.Error = err.Error()
	} else {
		res.Data = data
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("failed to render json: %v\n", err)

		msg := "internal error"
		msg = escapeJSON(msg)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, jsonErrTmpl, msg)
	}
}

// escapeJSON does primitive JSON escaping.
func escapeJSON(s string) string {
	return strings.Replace(s, `"`, `\"`, -1)
}

// jsonErrTmpl is the template to use when returning a JSON error. It is
// rendered using Printf, not json.Encode, so values must be escaped by
// the caller.
const jsonErrTmpl = `{"errorMsg":"%s"}`
