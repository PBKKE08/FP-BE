package repository

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/booking/order"
	"github.com/PBKKE08/FP-BE/core/model/booking/transaction"
	"github.com/PBKKE08/FP-BE/core/model/kategori"
	"github.com/PBKKE08/FP-BE/core/model/kota"
	"github.com/PBKKE08/FP-BE/core/model/review"

	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
)

type Pengguna interface {
	ByID(ctx context.Context, id pengguna.ID) (pengguna.Pengguna, error)
	Save(ctx context.Context, user pengguna.Pengguna) error
}

type Partner interface {
	ByID(ctx context.Context, id partner.ID) (partner.Partner, error)
	Save(ctx context.Context, user partner.Partner) error
	Approved(ctx context.Context, id partner.ID) error
	Delete(ctx context.Context, id partner.ID) error
}

type Kota interface {
	ByID(ctx context.Context, id kota.ID) (kota.Kota, error)
}

type Review interface {
	Save(ctx context.Context, review review.Review) error
}

type Transaction interface {
	Save(ctx context.Context, tx transaction.Transaction) error
}

type Order interface {
	ByID(ctx context.Context, id order.ID) (order.Order, error)
	Save(ctx context.Context, order order.Order) error
}

type Kategori interface {
	ByID(ctx context.Context, id kategori.ID) (kategori.Kategori, error)
}
