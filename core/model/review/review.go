package review

import (
	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
)

type Review struct {
	ID       ID
	Pengguna pengguna.Pengguna
	Partner  partner.Partner
	Rating   int
	Comment  string
}
