package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/purusah/cxtnxbr/pkg/config"
	"github.com/purusah/cxtnxbr/pkg/globals"
)

func DummyHandler(w http.ResponseWriter, r *http.Request) {}

func TestCounterMiddleware(t *testing.T) {
	url := "/api"
	conf, _ := config.GetConfig()
	gl := globals.NewCache(conf)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	CounterMiddleware(gl.C, http.HandlerFunc(DummyHandler)).ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Errorf("expected %d got %d code", http.StatusOK, res.Code)
	}
	val, _ := gl.C.Get(context.Background(), GetUriCounterKey(url))
	if val != 1 {
		t.Errorf("expected %d value, got %d", 1, val)
	}
}
