package braintree

import (
	"encoding/xml"
)

type Address struct {
	XMLName            xml.Name
	Id                 string `xml:"id,omitempty" json:"-" schema:"-"`
	CustomerId         string `xml:"customer-id,omitempty" json:"-" schema:"-"`
	FirstName          string `xml:"first-name,omitempty" json:"-" schema:"-"`
	LastName           string `xml:"last-name,omitempty" json:"-" schema:"-"`
	Company            string `xml:"company,omitempty" json:"-" schema:"-"`
	StreetAddress      string `xml:"street-address,omitempty" json:"street_address" schema:"street_address"`
	ExtendedAddress    string `xml:"extended-address,omitempty" json:"extended_address" schema:"extended_address"`
	Locality           string `xml:"locality,omitempty" json:"locality" schema:"locality"`
	Region             string `xml:"region,omitempty" json:"region" schema:"region"`
	PostalCode         string `xml:"postal-code,omitempty" json:"postal_code" schema:"postal_code"`
	CountryCodeAlpha2  string `xml:"country-code-alpha2,omitempty" json:"-" schema:"-"`
	CountryCodeAlpha3  string `xml:"country-code-alpha3,omitempty" json:"-" schema:"-"`
	CountryCodeNumeric string `xml:"country-code-numeric,omitempty" json:"-" schema:"-"`
	CountryName        string `xml:"country-name,omitempty" json:"-" schema:"-"`
	CreatedAt          string `xml:"created-at,omitempty" json:"-" schema:"-"`
	UpdatedAt          string `xml:"updated-at,omitempty" json:"-" schema:"-"`
}

type Addresses struct {
	XMLName string     `xml:"addresses"`
	Address []*Address `xml:"address"`
}
