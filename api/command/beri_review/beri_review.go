package command

import (
	"context"

	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
	"github.com/PBKKE08/FP-BE/core/model/review"
)

type BeriReview struct {
	PenggunaRepo PenggunaRepository
	PartnerRepo  PartnerRepository
	ReviewRepo   ReviewRepository
}

func (b *BeriReview) Execute(ctx context.Context, req BeriReviewRequest) error {
	penggunaId, err := pengguna.NewIDFrom(req.PenggunaID)
	if err != nil {
		return err
	}

	pengguna, err := b.PenggunaRepo.ByID(ctx, penggunaId)
	if err != nil {
		return err
	}

	partnerId, err := partner.NewIDFrom(req.PartnerID)
	if err != nil {
		return err
	}

	partner, err := b.PartnerRepo.ByID(ctx, partnerId)
	if err != nil {
		return err
	}

	reviewID := review.NewID()
	reviewInsert := review.Review{
		ID:       reviewID,
		Pengguna: pengguna,
		Partner:  partner,
		Rating:   req.Rating,
		Comment:  req.Comment,
	}

	err = b.ReviewRepo.Save(ctx, reviewInsert)
	if err != nil {
		return err
	}

	return nil
}
