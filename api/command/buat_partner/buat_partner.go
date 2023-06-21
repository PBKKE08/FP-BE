package buat_partner

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/kategori"
	"github.com/PBKKE08/FP-BE/core/model/kota"
	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/repository"
)

type BuatPartner struct {
	PartnerRepo  repository.Partner
	KotaRepo     repository.Kota
	KategoriRepo repository.Kategori
}

func (u *BuatPartner) Execute(ctx context.Context, req Request) error {
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

	partnerData := partner.Partner{
		ID:           partner.NewID(),
		Nama:         req.Nama,
		Email:        req.Email,
		NomorTelepon: req.NomorTelepon,
		JenisKelamin: req.JenisKelamin,
		Kota:         kotaData,
		Harga:        req.Harga,
		Kategori:     kategoriData,
		Description:  req.Description,
	}

	err = u.PartnerRepo.Save(ctx, partnerData)

	return err
}
