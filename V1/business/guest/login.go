package guest

import (
	"github.com/PBKKE08/FP-BE/V1/business/common"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	common.Response
}

func (s *Service) Login(in LoginInput) LoginOutput {
	var out LoginOutput

	user, err := s.repo.GetByEmail(in.Email)
	if err != nil {
		out.SetError(500, err)
		return out
	}

	if user == emptyUser {
		out.SetError(400, ErrUnknownUser)
		return out
	}

	if !passwordMatch(user.Password, in.Password) {
		out.SetError(400, ErrPasswordNotMatch)
		return out
	}

	token, err := s.tokenGenerator.GenerateToken(user.ID.String(), user.Name, user.Email, user.Telephone, user.Gender)
	if err != nil {
		out.SetError(500, err)
		return out
	}

	out.Token = token
	out.SetOK()

	return out
}

func passwordMatch(userPassword, incomingPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(incomingPassword))
	return err == nil
}
