package partner

import "github.com/PBKKE08/FP-BE/V1/business/common"

type AllOrderInput struct {
	PartnerID string
}

type AllOrderOutput struct {
	common.Response
}

func (s *Service) GetAllOrder(in AllOrderInput) AllOrderOutput {
	var out AllOrderOutput

	return out
}
