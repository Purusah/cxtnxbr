package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Counter interface {
	Incr(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (int, error)
}

func GetUriCounterKey(uri string) string {
	return fmt.Sprintf("counter:%s", uri)
}

func CounterMiddleware(c Counter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := c.Incr(r.Context(), GetUriCounterKey(r.RequestURI))
		next.ServeHTTP(w, r)
		if err != nil {
			log.Printf("counter mw inc %e", err)
		}
	})
}
