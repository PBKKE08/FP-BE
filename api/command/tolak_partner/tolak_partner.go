package tolak_partner

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/repository"
)

type TolakPartner struct {
	PartnerRepo repository.Partner
}

func (t *TolakPartner) Execute(ctx context.Context, id string) error {
	partnerID, err := partner.NewIDFrom(id)
	if err != nil {
		return err
	}

	return t.PartnerRepo.Delete(ctx, partnerID)
}
