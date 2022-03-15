package locations_provider

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
	rest.FlushMockups() // Works like a teardown before start the mocks
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CA",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found","error": "not_found","status": 404,"cause": []}`,
	})

	country, err := GetCountry("CA")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryRestclientError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CA",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})

	country, err := GetCountry("CA")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid restclient response. Cannot get the country CA", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CA",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found","error": "not_found","status": "404","cause": []}`,
	})

	country, err := GetCountry("CA")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface when getting country CA", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CA",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 123,"name": "Argentina","time_zone": "GMT-03:00"}`,
	})

	country, err := GetCountry("CA")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Error trying to unmarshall country data for CA", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CA",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"CA","name":"Canada","locale":"en_US","currency_id":"USD","decimal_separator":".","thousands_separator":",","time_zone":"GMT-04:00","geo_information":{"location":{"latitude":50.8286074,"longitude":-130.1191663}},"states":[{"id":"CA-AB","name":"Alberta"},{"id":"CA-BC","name":"British Columbia"},{"id":"CA-MB","name":"Manitoba"},{"id":"CA-NB","name":"New Brunswick"},{"id":"CA-NL","name":"Newfoundland and Lab"},{"id":"CA-NT","name":"Northwest Territorie"},{"id":"CA-NS","name":"Nova Scotia"},{"id":"CA-NU","name":"Nunavut"},{"id":"CA-ON","name":"Ontario"},{"id":"CA-PE","name":"Prince Edward Island"},{"id":"CA-QC","name":"Quebec"},{"id":"CA-YT","name":"Yukon"}]}`,
	})

	country, err := GetCountry("CA")

	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "CA", country.Id)
	assert.EqualValues(t, "Canada", country.Name)
	assert.EqualValues(t, "GMT-04:00", country.TimeZone)
	assert.EqualValues(t, 12, len(country.States))
}
