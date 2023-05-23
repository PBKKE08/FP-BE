package common

import "github.com/nyaruka/phonenumbers"

type TelephoneNumber struct {
	internationalFormatted string
	nationalFormatted      string
}

var emptyTelephoneNumber = TelephoneNumber{}

func NewTelephoneNumber(from string) (TelephoneNumber, error) {
	num, err := phonenumbers.Parse(from, "ID")
	if err != nil {
		return emptyTelephoneNumber, err
	}

	var t TelephoneNumber
	t.nationalFormatted = phonenumbers.Format(num, phonenumbers.NATIONAL)
	t.internationalFormatted = phonenumbers.Format(num, phonenumbers.INTERNATIONAL)

	return t, nil
}

func (t TelephoneNumber) ToNationalFormat() string {
	return t.nationalFormatted
}

func (t TelephoneNumber) ToInternationalFormat() string {
	return t.internationalFormatted
}
