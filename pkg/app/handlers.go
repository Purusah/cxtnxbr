package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/purusah/cxtnxbr/pkg/globals"
	"github.com/purusah/cxtnxbr/pkg/middlewares"
)

type WsCounterRequest struct {
	Route string `json:"route"`
}

type WsCounterResponse struct {
	Route  string `json:"route"`
	Amount int    `json:"amount"`
}

type WsErrorResponse struct {
	Error string `json:"error"`
}

func WsCounterHandler(gl *globals.Global) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req WsCounterRequest
		c, err := gl.WsUpgrade.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("ws upgrade %s", err.Error())
			return
		}
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("read ws %s", err.Error())
				break
			}
			err = json.Unmarshal(message, &req)
			if err != nil {
				errMsg, _ := json.Marshal(WsErrorResponse{Error: "error json reading"})
				err = c.WriteMessage(mt, errMsg)
				if err != nil {
					log.Printf("write ws %s", err.Error())
					break
				}
				continue
			}
			amount, err := gl.C.Get(r.Context(), middlewares.GetUriCounterKey(req.Route))
			if err != nil {
				log.Printf("cache error get %s", err.Error())
				errMsg, _ := json.Marshal(WsErrorResponse{Error: "error data reading"})
				err = c.WriteMessage(mt, errMsg)
				if err != nil {
					log.Printf("write ws %s", err.Error())
					break
				}
				continue
			}
			resSerialized, _ := json.Marshal(WsCounterResponse{Route: req.Route, Amount: amount})
			err = c.WriteMessage(mt, resSerialized)
			if err != nil {
				log.Printf("write ws %s", err.Error())
				break
			}
		}
	}
}

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
