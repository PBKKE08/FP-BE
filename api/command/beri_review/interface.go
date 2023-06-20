package command

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/review"

	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
)

type PenggunaRepository interface {
	ByID(ctx context.Context, id pengguna.ID) (pengguna.Pengguna, error)
}

type PartnerRepository interface {
	ByID(ctx context.Context, id partner.ID) (partner.Partner, error)
}

type ReviewRepository interface {
	Save(ctx context.Context, review review.Review) error
}
