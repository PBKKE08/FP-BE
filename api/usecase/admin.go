package usecase

import (
	"context"
	"fmt"
	"github.com/PBKKE08/FP-BE/api/command/terima_pembayaran"
	"github.com/PBKKE08/FP-BE/api/query"
)

type TolakPartnerCommand interface {
	Execute(ctx context.Context, id string) error
}

type TerimaPartnerCommand interface {
	Execute(ctx context.Context, id string) error
}

type TerimaPembayaranCommand interface {
	Execute(ctx context.Context, r terima_pembayaran.Request) error
}

type DaftarListPendaftarQuery interface {
	GetListPendaftar(ctx context.Context) []query.PartnerInginMendaftar
}

type TransaksiNonConfirmedQuery interface {
	DaftarTransaksiNonConfirmed(ctx context.Context) []query.TransaksiNonTerbayar
}

type AdminAuthProvider interface {
	ApprovedAcc(ctx context.Context, email string) error
	DeleteAcc(ctx context.Context, email string) error
}

type AdminUsecase struct {
	getListPendaftar      DaftarListPendaftarQuery
	tolakPartner          TolakPartnerCommand
	terimaPartner         TerimaPartnerCommand
	transaksiNonConfirmed TransaksiNonConfirmedQuery
	terimaPembayaran      TerimaPembayaranCommand
	mailer                Mailer
	authProvider          AdminAuthProvider
}

func NewAdminUsecase(
	getListPendaftar DaftarListPendaftarQuery,
	tolakPartner TolakPartnerCommand,
	terimaPartner TerimaPartnerCommand,
	confirmedQuery TransaksiNonConfirmedQuery,
	terimaPembayaran TerimaPembayaranCommand,
	mailer Mailer,
	authProvider AdminAuthProvider) *AdminUsecase {
	return &AdminUsecase{
		getListPendaftar:      getListPendaftar,
		tolakPartner:          tolakPartner,
		terimaPartner:         terimaPartner,
		authProvider:          authProvider,
		terimaPembayaran:      terimaPembayaran,
		mailer:                mailer,
		transaksiNonConfirmed: confirmedQuery}
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

func (a *AdminUsecase) DaftarTxNonConfirmed(ctx context.Context) []query.TransaksiNonTerbayar {
	return a.transaksiNonConfirmed.DaftarTransaksiNonConfirmed(ctx)
}

func (a *AdminUsecase) TerimaPembayaranPengguna(ctx context.Context, r terima_pembayaran.Request) error {
	err := a.terimaPembayaran.Execute(ctx, r)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("From: %s | Date: %s | Start: %s | End: %s", r.NamaPembooking, r.WaktuBooking, r.Mulai, r.Selesai)

	err = a.mailer.Mail(ctx, "socium@company.com", r.Email, "New Order", msg)
	if err != nil {
		return err
	}

	return nil
}
