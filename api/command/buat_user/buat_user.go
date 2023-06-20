package buat_user

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/kota"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
	"github.com/PBKKE08/FP-BE/core/repository"
)

type BuatUser struct {
	PenggunaRepo repository.Pengguna
	KotaRepo     repository.Kota
}

func (u *BuatUser) Execute(ctx context.Context, req BuatUserRequest) error {
	kotaID, err := kota.NewIDFrom(req.KotaID)
	if err != nil {
		return err
	}

	kotaData, err := u.KotaRepo.ByID(ctx, kotaID)
	if err != nil {
		return err
	}

	userID := pengguna.NewID()

	user := pengguna.Pengguna{
		ID:           userID,
		Nama:         req.Nama,
		Email:        req.Email,
		NomorTelepon: req.Telepon,
		JenisKelamin: req.JenisKelamin,
		Kota:         kotaData,
	}

	return u.PenggunaRepo.Save(ctx, user)
}
