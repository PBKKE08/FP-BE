package usecase

import (
	"context"
	"errors"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"

	command "github.com/PBKKE08/FP-BE/api/command/beri_review"
	"github.com/PBKKE08/FP-BE/api/query"
)

var (
	ErrUnknownJenisKelamin = errors.New("jenis kelamin tidak diketahui")
	ErrRatingMelebihiBatas = errors.New("rating maksimal hanya 5")
)

type CariPasanganQuery interface {
	By(ctx context.Context, daerah string, jenisKelamin string, kebutuhan string) []query.CariPasangan
}

type LihatTransaksiQuery interface {
	LihatTransaksi(ctx context.Context, id pengguna.ID) ([]query.SeluruhTransaksi, error)
}

type BeriReviewCommand interface {
	Execute(ctx context.Context, req command.BeriReviewRequest) error
}

type PenggunaUsecase struct {
	cariPasanganQuery CariPasanganQuery
	buatReviewCommand BeriReviewCommand
	lihatTransaksi    LihatTransaksiQuery
}

func NewPenggunaUsecase(cariPasanganQuery CariPasanganQuery, buatReviewCommand BeriReviewCommand, lihatTransaksi LihatTransaksiQuery) *PenggunaUsecase {
	return &PenggunaUsecase{
		cariPasanganQuery: cariPasanganQuery,
		buatReviewCommand: buatReviewCommand,
		lihatTransaksi:    lihatTransaksi,
	}
}

func (p *PenggunaUsecase) CariPasanganBerdasarkan(ctx context.Context, daerah string, jenisKelamin string, kebutuhan string) ([]query.CariPasangan, error) {
	if jenisKelamin != "" && !(jenisKelamin == "f" || jenisKelamin == "m") {
		return []query.CariPasangan{}, ErrUnknownJenisKelamin
	}

	partners := p.cariPasanganQuery.By(ctx, daerah, jenisKelamin, kebutuhan)

	return partners, nil
}

func (p *PenggunaUsecase) LihatRiwayaTransaksi(ctx context.Context, id string) ([]query.SeluruhTransaksi, error) {
	penggunaID, err := pengguna.NewIDFrom(id)
	if err != nil {
		return nil, err
	}

	results, err := p.lihatTransaksi.LihatTransaksi(ctx, penggunaID)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (p *PenggunaUsecase) ReviewPartner(ctx context.Context, req command.BeriReviewRequest) error {
	if req.Rating > 5 {
		return ErrRatingMelebihiBatas
	}

	err := p.buatReviewCommand.Execute(ctx, req)
	return err
}
