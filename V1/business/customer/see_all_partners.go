package customer

import "github.com/PBKKE08/FP-BE/V1/business/common"

type AllPartnerInput struct {
	FilterByGender   string `json:"-"`
	FilterByName     string `json:"-"`
	FilterByCategory string `json:"-"`
}

type AllPartnerOutput struct {
	common.Response
	Partners []Partner `json:"partners,omitempty"`
}

func (s *Service) SeeAllPartners(in AllPartnerInput) AllPartnerOutput {
	var out AllPartnerOutput

	queryParam := mapQueryParam(in.FilterByGender, in.FilterByName, in.FilterByCategory)

	partners, err := s.repo.SeeAllPartners(queryParam)
	if err != nil {
		out.SetError(500, err)
		return out
	}

	out.Partners = partners

	out.SetOK()
	return out
}

func mapQueryParam(queryByGender, queryByName, queryByCategory string) map[string]string {
	m := make(map[string]string, 3)

	if queryByGender != "" {
		m["gender"] = queryByGender
	}

	if queryByName != "" {
		m["name"] = queryByName
	}

	if queryByCategory != "" {
		m["category"] = queryByCategory
	}

	return m
}
