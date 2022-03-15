package services

import (
	"testing-golang/src/api/domain/locations"
	"testing-golang/src/api/utils/errors"
	"testing-golang/src/providers/locations_provider"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
