package guest

import (
	"github.com/PBKKE08/FP-BE/V1/business/common"
	"io"
)

type RegisterPartnerInput struct {
	Email     string             `json:"email" form:"email"`
	Password  string             `json:"password" form:"password"`
	Name      string             `json:"nama" form:"nama""`
	Telephone string             `json:"nomor_telepon" form:"nomor_telepon"`
	Gender    string             `json:"gender" form:"gender"`
	Age       string             `json:"umur"`
	Image     io.ReadWriteCloser `form:"wajah"`
	Reason    string             `json:"alasan" form:"alasan"`
	Harga     string             `json:"harga" form:"harga"`
}

type RegisterPartnerOutput struct {
	common.Response
}

func (s *Service) RegisterPartner(in RegisterInput) RegisterPartnerOutput {
	var out RegisterPartnerOutput

	user, err := s.repo.GetByEmail(in.Email)
	if err != nil {
		out.SetError(500, err)
		return out
	}

	if user != emptyUser {
		out.SetError(400, ErrEmailAlreadyExists)
		return out
	}

	hashedPassword := hashPassword(in.Password)

	imgURL, err := s.imageSaver.SaveImage(in.Image)
	if err != nil {
		out.SetError(500, err)
		return out
	}

	err = s.repo.InsertUser(in.Name, in.Email, in.Telephone, hashedPassword, in.Gender, in.Age, imgURL, RolePartner)
	if err != nil {
		out.SetError(500, err)
		return out
	}

	GuestWasRegistered(s.mailer, user)

	out.Set(201, "Registered")
	return out
}
