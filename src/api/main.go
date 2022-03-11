package main

import (
	"fmt"
	"testing-golang/src/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("CA")

	fmt.Println(err)
	fmt.Println(country)
}
