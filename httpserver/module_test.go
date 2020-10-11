package httpserver_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"

	"github.com/lucaskatayama/fx-contrib/httpserver"
)

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
		return
	})
	return mux
}

func TestSimple(t *testing.T) {

	app := fxtest.New(t)
	app.App = fx.New(httpserver.Module, fx.Provide(handler))
	defer app.RequireStop()
	app.RequireStart()

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/test", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.FailNow()
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.FailNow()
	}
	if string(b) != "Hello" {
		t.FailNow()
	}

}
