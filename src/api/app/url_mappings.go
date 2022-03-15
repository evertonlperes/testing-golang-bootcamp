package app

import (
	"testing-golang/src/api/controllers"
)

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
