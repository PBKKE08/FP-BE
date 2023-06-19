package usecase

import (
	"context"
	"errors"

	command "github.com/PBKKE08/FP-BE/api/command/beri_review"
	"github.com/PBKKE08/FP-BE/api/query"
)

var (
	ErrUnknownJenisKelamin = errors.New("jenis kelamin tidak diketahui")
	ErrRatingMelebihiBatas = errors.New("rating maksimal hanya 5")
)

type PenggunaUsecase struct {
	cariPasanganQuery CariPasanganQuery
	buatReviewCommand BeriReviewCommand
}

func NewPenggunaUsecase(cariPasanganQuery CariPasanganQuery, buatReviewCommand BeriReviewCommand) *PenggunaUsecase {
	return &PenggunaUsecase{
		cariPasanganQuery: cariPasanganQuery,
		buatReviewCommand: buatReviewCommand,
	}
}

func (p *PenggunaUsecase) CariPasanganBerdasarkan(ctx context.Context, daerah string, jenisKelamin string) ([]query.CariPasangan, error) {
	if !(jenisKelamin == "f" || jenisKelamin == "m") {
		return []query.CariPasangan{}, ErrUnknownJenisKelamin
	}

	partners := p.cariPasanganQuery.By(ctx, daerah, jenisKelamin)

	return partners, nil
}

func (p *PenggunaUsecase) ReviewPartner(ctx context.Context, req command.BeriReviewRequest) error {
	if req.Rating > 5 {
		return ErrRatingMelebihiBatas
	}

	err := p.buatReviewCommand.Execute(ctx, req)
	return err
}