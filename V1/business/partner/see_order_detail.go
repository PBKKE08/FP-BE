package partner

import "github.com/PBKKE08/FP-BE/V1/business/common"

type OrderDetailInput struct {
	PartnerID string
}

type OrderDetailOutput struct {
	common.Response
}

func (s *Service) GetOrderDetail(in OrderDetailInput) OrderDetailOutput {
	var out OrderDetailOutput

	return out
}
