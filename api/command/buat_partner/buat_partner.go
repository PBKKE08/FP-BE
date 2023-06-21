package buat_partner

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/kategori"
	"github.com/PBKKE08/FP-BE/core/model/kota"
	"github.com/PBKKE08/FP-BE/core/repository"
)

type BuatUser struct {
	PartnerRepo repository.Partner
	KotaRepo     repository.Kota
	KategoriRepo repository.Kategori
}

func (u *BuatUser) Execute(ctx context.Context, req Request) error {
	kotaID, err := kota.NewIDFrom(req.KotaID)
	if err != nil {
		return err
	}

	kotaData, err := u.KotaRepo.ByID(ctx, kotaID)
	if err != nil {
		return err
	}

	kategoriID, err := kategori.NewIDFrom(req.KategoriID)
	if err != nil {
		return err
	}

	kategoriData, err := u.KategoriRepo.ByID(ctx, kategoriID)
	if err != nil {
		return err
	}

	err = u.PartnerRepo.
}
