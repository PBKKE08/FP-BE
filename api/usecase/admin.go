package usecase

import (
	"context"
	"github.com/PBKKE08/FP-BE/api/query"
)

type TolakPartnerCommand interface {
	Execute(ctx context.Context, id string) error
}

type TerimaPartnerCommand interface {
	Execute(ctx context.Context, id string) error
}

type DaftarListPendaftarQuery interface {
	GetListPendaftar(ctx context.Context) []query.PartnerInginMendaftar
}

type AdminAuthProvider interface {
	ApprovedAcc(ctx context.Context, email string) error
	DeleteAcc(ctx context.Context, email string) error
}

type AdminUsecase struct {
	getListPendaftar DaftarListPendaftarQuery
	tolakPartner     TolakPartnerCommand
	terimaPartner    TerimaPartnerCommand
	authProvider     AdminAuthProvider
}

func NewAdminUsecase(
	getListPendaftar DaftarListPendaftarQuery,
	tolakPartner TolakPartnerCommand,
	terimaPartner TerimaPartnerCommand,
	authProvider AdminAuthProvider) *AdminUsecase {
	return &AdminUsecase{getListPendaftar: getListPendaftar, tolakPartner: tolakPartner, terimaPartner: terimaPartner, authProvider: authProvider}
}

func (a *AdminUsecase) GetPartnerYangInginMendaftar(ctx context.Context) []query.PartnerInginMendaftar {
	return a.getListPendaftar.GetListPendaftar(ctx)
}

func (a *AdminUsecase) TerimaPartnerPendaftar(ctx context.Context, id, email string) error {
	err := a.terimaPartner.Execute(ctx, id)
	if err != nil {
		return err
	}

	err = a.authProvider.ApprovedAcc(ctx, email)
	if err != nil {
		return err
	}

	return nil
}

func (a *AdminUsecase) TolakPartnerPendaftar(ctx context.Context, id string, email string) error {
	err := a.tolakPartner.Execute(ctx, id)
	if err != nil {
		return err
	}

	err = a.authProvider.DeleteAcc(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
