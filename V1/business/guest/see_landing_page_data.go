package guest

import "github.com/PBKKE08/FP-BE/V1/business/common"

type LandingPageDataOutput struct {
	common.Response
	BestPartner       []Partner `json:"best_partner"`
	PopularCategories []string  `json:"popular_categories"`
}

func (s *Service) SeeLandingPageData() LandingPageDataOutput {
	var out LandingPageDataOutput

	landingPageData, err := s.repo.GetLandingPageData()
	if err != nil {
		out.SetError(500, err)
		return out
	}

	out.BestPartner = landingPageData.BestPartner
	out.PopularCategories = landingPageData.PopularCategories

	out.SetOK()
	return out
}
