package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"testing-golang/src/api/utils/errors"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesNotFound(t *testing.T) {
	// TODO: Need to be fixed, getting interface error instead notFound
	//Init: Mock the server
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CA",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"status": 404, "error": "not_found", "message": "no country with id CA"`,
	})
	// Execution: Hit the api and get the response
	response, err := http.Get("http://localhost:8080/locations/countries/CA")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	bytes, _ := ioutil.ReadAll(response.Body)

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	fmt.Println(err)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "no country with id CA", apiErr.Message)
}

func TestGetCountriesNoError(t *testing.T) {
	//TODO: Implement the test scenario.
}
