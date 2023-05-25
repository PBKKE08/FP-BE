package customer

import (
	"errors"
	"github.com/PBKKE08/FP-BE/V1/business/common"
)

var (
	unknownPartner    = Partner{}
	ErrUnknownPartner = errors.New("unknown partner")
)

type GetPartnerDetailInput struct {
	ID string
}

type GetPartnerDetailOutput struct {
	common.Response
	Partner Partner `json:"partner"`
}

func (s *Service) GetPartnerDetail(in GetPartnerDetailInput) GetPartnerDetailOutput {
	var out GetPartnerDetailOutput

	partnerDetail, err := s.repo.GetPartnerDetail(in.ID)
	if err != nil {
		out.SetError(500, err)
		return out
	}

	if partnerDetail == unknownPartner {
		out.SetError(400, ErrUnknownPartner)
		return out
	}

	out.Partner = partnerDetail
	out.SetOK()

	return out
}
