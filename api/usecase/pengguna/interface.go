package usecase

import (
	"context"

	command "github.com/PBKKE08/FP-BE/api/command/beri_review"
	"github.com/PBKKE08/FP-BE/api/query"
)

type CariPasanganQuery interface {
	By(ctx context.Context, daerah string, jenisKelamin string) []query.CariPasangan
}

type BeriReviewCommand interface {
	Execute(ctx context.Context, req command.BeriReviewRequest) error
}
