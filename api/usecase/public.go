package usecase

import (
	"context"
	"github.com/PBKKE08/FP-BE/api/query"
)

type CitiesCategoriesQuery interface {
	GetAllCityAndCategory(ctx context.Context) query.AllCitiesAndCategories
}

type PublicUsecase struct {
	getter CitiesCategoriesQuery
}

func NewPublicUsecase(getter CitiesCategoriesQuery) *PublicUsecase {
	return &PublicUsecase{getter: getter}
}

func (p *PublicUsecase) GetAllCityAndCategory(ctx context.Context) query.AllCitiesAndCategories {
	return p.getter.GetAllCityAndCategory(ctx)
}
