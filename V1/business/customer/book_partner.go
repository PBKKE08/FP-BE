package customer

import (
	"errors"
	"github.com/PBKKE08/FP-BE/V1/business/common"
	"time"
)

var (
	ErrPartnerHasBeenBookedAtThisTime = errors.New("partner has been booked in this time, choose another one")
)

type BookingPartnerInput struct {
	PartnerID             string    `json:"partner_id"`
	City                  string    `json:"city"`
	Date                  time.Time `json:"date"`
	From                  time.Time `json:"from"`
	To                    time.Time `json:"to"`
	Category              string    `json:"category"`
	AdditionalInformation string    `json:"additional_information"`

	UserEmail string `json:"-"`
	UserName  string `json:"-"`
}

type BookingPartnerOutput struct {
	common.Response
	Partner Partner `json:"partner"`
	Cost    string  `json:"cost"`
}

func (s *Service) BookPartner(in BookingPartnerInput) BookingPartnerOutput {
	var out BookingPartnerOutput

	partnerDetail, err := s.repo.GetPartnerDetail(in.PartnerID)
	if err != nil {
		out.SetError(500, err)
		return out
	}

	if partnerDetail == unknownPartner {
		out.SetError(400, ErrUnknownPartner)
		return out
	}

	if !isPartnerAvailableAt(in.Date, in.From, in.To) {
		out.SetError(400, ErrPartnerHasBeenBookedAtThisTime)
		return out
	}

	duration := in.To.Sub(in.From)
	customerTotalCost := getCustomerTotalCost(partnerDetail.Price, duration)

	PartnerWasBooked(in.UserName, in.UserEmail, partnerDetail, customerTotalCost)

	out.Cost = customerTotalCost
	out.Partner = partnerDetail
	out.SetOK()

	return out
}

// TODO
func isPartnerAvailableAt(date, from, to time.Time) bool {
	return true
}

// TODO
func getCustomerTotalCost(price string, difference time.Duration) string {
	return ""
}
