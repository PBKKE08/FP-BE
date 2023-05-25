package guest

import (
	"errors"
	"github.com/PBKKE08/FP-BE/V1/business/common"
	"github.com/PBKKE08/FP-BE/V1/entity"
	"golang.org/x/crypto/bcrypt"
	"io"
	"unsafe"
)

var (
	ErrPasswordNotMatch   = errors.New("bad credentials")
	ErrEmailAlreadyExists = errors.New("email already taken")
	ErrUnknownUser        = errors.New("unknown user")
)

var (
	emptyUser = entity.User{}
)

const (
	RoleCustomer = "customer"
	RolePartner  = "partner"
)

type RegisterInput struct {
	Email     string             `json:"email" form:"email"`
	Password  string             `json:"password" form:"password"`
	Name      string             `json:"nama" form:"nama""`
	Telephone string             `json:"nomor_telepon" form:"nomor_telepon"`
	Gender    string             `json:"gender" form:"gender"`
	Age       string             `json:"umur"`
	Image     io.ReadWriteCloser `form:"wajah"`
}

type RegisterOutput struct {
	common.Response
}

func (s *Service) Register(in RegisterInput) RegisterOutput {
	var out RegisterOutput

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

	err = s.repo.InsertUser(in.Name, in.Email, in.Telephone, hashedPassword, in.Gender, in.Age, imgURL, RoleCustomer)
	if err != nil {
		out.SetError(500, err)
		return out
	}

	GuestWasRegistered(s.mailer, user)

	out.Set(201, "Registered")
	return out
}

func hashPassword(password string) string {
	s, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return *(*string)(unsafe.Pointer(&s))
}
