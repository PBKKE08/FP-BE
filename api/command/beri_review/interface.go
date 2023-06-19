package command

import (
	"context"

	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
)

type PenggunaRepository interface {
	Exist(ctx context.Context, id pengguna.ID) error
}

type PartnerRepository interface {
	Exist(ctx context.Context, id partner.ID) error
}

type ReviewRepository interface {
	Save(ctx context.Context, penggunaID pengguna.ID, partnerID partner.ID, rating int, comment string) error
}
