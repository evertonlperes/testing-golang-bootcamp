package test

import (
	"os"
	"testing"
	"testing-golang/src/api/app"

	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp() //need to start the app in another goRoutine
	os.Exit(m.Run())
}
