package repository

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/kota"
	"github.com/PBKKE08/FP-BE/core/model/review"

	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
)

type Pengguna interface {
	ByID(ctx context.Context, id pengguna.ID) (pengguna.Pengguna, error)
	Save(ctx context.Context, user pengguna.Pengguna) error
}

type Partner interface {
	ByID(ctx context.Context, id partner.ID) (partner.Partner, error)
}

type Kota interface {
	ByID(ctx context.Context, id kota.ID) (kota.Kota, error)
}

type Review interface {
	Save(ctx context.Context, review review.Review) error
}
