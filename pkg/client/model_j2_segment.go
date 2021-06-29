/*
 * METRO2 API
 *
 * Moov Metro2 ([Automated Clearing House](https://en.wikipedia.org/wiki/Automated_Clearing_House)) implements an HTTP API for creating, parsing and validating Metro2 files. Metro2 is an open-source consumer credit history report for credit report file creation and validation.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"time"
)

// J2Segment struct for J2Segment
type J2Segment struct {
	SegmentIdentifier            string    `json:"segmentIdentifier"`
	Surname                      string    `json:"surname"`
	FirstName                    string    `json:"firstName"`
	MiddleName                   string    `json:"middleName,omitempty"`
	GenerationCode               string    `json:"generationCode,omitempty"`
	SocialSecurityNumber         int32     `json:"socialSecurityNumber,omitempty"`
	DateBirth                    time.Time `json:"dateBirth"`
	TelephoneNumber              int64     `json:"telephoneNumber,omitempty"`
	EcoaCode                     string    `json:"ecoaCode"`
	ConsumerInformationIndicator string    `json:"consumerInformationIndicator,omitempty"`
	CountryCode                  string    `json:"countryCode,omitempty"`
	FirstLineAddress             string    `json:"firstLineAddress"`
	SecondLineAddress            string    `json:"secondLineAddress,omitempty"`
	City                         string    `json:"city"`
	State                        string    `json:"state"`
	ZipCode                      string    `json:"zipCode"`
	AddressIndicator             string    `json:"addressIndicator,omitempty"`
	ResidenceCode                string    `json:"residenceCode,omitempty"`
}
