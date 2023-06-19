package command

import (
	"context"

	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
)

type BeriReview struct {
	PenggunaRepo PenggunaRepository
	PartnerRepo PartnerRepository
	ReviewRepo ReviewRepository
}

func (b *BeriReview) Execute(ctx context.Context, req BeriReviewRequest) error {
	penggunaId, err := pengguna.NewIDFrom(req.PenggunaID)
	if err != nil {
		return err
	}
	
	err = b.PenggunaRepo.Exist(ctx, penggunaId)
	if err != nil {
		return err
	}

	partnerId, err := partner.NewIDFrom(req.PartnerID)
	if err != nil {
		return err
	}

	err = b.PartnerRepo.Exist(ctx, partnerId)
	if err != nil {
		return err
	}

	err = b.ReviewRepo.Save(ctx, penggunaId, partnerId, req.Rating, req.Comment)
	if err != nil {
		return err
	}

	return nil
}
