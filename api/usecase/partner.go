package usecase

import "context"

type LihatDetailDiriQuery interface {
	Execute(ctx context.Context)
}

type PartnerUsecase struct {
}
