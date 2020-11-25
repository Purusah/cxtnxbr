package app

import (
	"log"
	"net/http"

	"github.com/purusah/cxtnxbr/pkg/config"
	gl "github.com/purusah/cxtnxbr/pkg/globals"
	mw "github.com/purusah/cxtnxbr/pkg/middlewares"
)

func StartApp() {
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	env := gl.NewCache(conf)
	srv := http.NewServeMux()
	srv.HandleFunc("/ws/api/v1/counter", WsCounterHandler(env))
	srv.Handle("/api/v1/any", mw.CounterMiddleware(env.C, http.HandlerFunc(DefaultHandler(env))))

	log.Printf("start app %s", conf.App.Url)
	if err = http.ListenAndServe(conf.App.Url, srv); err != nil {
		log.Fatal(err)
	}
}
