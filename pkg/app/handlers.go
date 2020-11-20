package app

import (
	"log"
	"net/http"

	"github.com/purusah/cxtnxbr/pkg/globals"
	"github.com/purusah/cxtnxbr/pkg/middlewares"
)

func DefaultHandler(gl *globals.Global) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			amount, err := gl.C.Get(r.Context(), middlewares.GetUriCounterKey(r.RequestURI))
			if err != nil {
				log.Print("cache error get ", err)
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			err = gl.T.Counter.Execute(w, struct {
				Amount int
			}{Amount: amount})
			if err != nil {
				log.Print("template exec ", err)
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		default:
			http.Error(w, "", http.StatusMethodNotAllowed)
		}
	}
}
