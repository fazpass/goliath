package phonenumber

import (
	"fmt"
	"strings"
)

type Country struct {
	ISOCode     string
	CountryCode string
	Name        string
}

var countries = []Country{
	{
		ISOCode:     "AUS",
		CountryCode: "61",
		Name:        "Australia",
	},
	{
		ISOCode:     "BRA",
		CountryCode: "55",
		Name:        "Brazil",
	},
	{
		ISOCode:     "IDN",
		CountryCode: "62",
		Name:        "Indonesia",
	},
	{
		ISOCode:     "GBR",
		CountryCode: "44",
		Name:        "United Kingdom",
	},
	{
		ISOCode:     "AUT",
		CountryCode: "43",
		Name:        "Austria",
	},
	{
		ISOCode:     "SGP",
		CountryCode: "65",
		Name:        "Singapore",
	},
	{
		ISOCode:     "USA",
		CountryCode: "1",
		Name:        "United States",
	},
}

func GetCountryByPhoneNumber(phoneNumber string) (*Country, error) {
	phoneNumber = strings.Replace(phoneNumber, "+", "", 1)

	for _, country := range countries {
		countryCodeLength := len(country.CountryCode)
		phoneNumberCountryCode := phoneNumber[0:countryCodeLength]

		if country.CountryCode == phoneNumberCountryCode {
			return &country, nil
		}
	}

	return nil, fmt.Errorf("phone number not recognized")
}
