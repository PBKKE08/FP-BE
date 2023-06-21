package usecase

import (
	"context"
	"github.com/PBKKE08/FP-BE/api/query"
	"github.com/PBKKE08/FP-BE/core/model/partner"
)

type DetailPartnerQuery interface {
	LihatPartnerDetail(ctx context.Context, id partner.ID) (query.DetailPartner, error)
}

type PartnerUsecase struct {
	getPartnerDetail DetailPartnerQuery
}

func NewPartnerUsecase(getPartnerDetail DetailPartnerQuery) *PartnerUsecase {
	return &PartnerUsecase{getPartnerDetail: getPartnerDetail}
}

func (p *PartnerUsecase) GetPartnerDetail(ctx context.Context, id string) (query.DetailPartner, error) {
	partnerID, err := partner.NewIDFrom(id)
	if err != nil {
		return query.DetailPartner{}, err
	}

	return p.getPartnerDetail.LihatPartnerDetail(ctx, partnerID)
}
