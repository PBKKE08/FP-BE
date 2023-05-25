package guest

import (
	"github.com/PBKKE08/FP-BE/V1/entity"
	"io"
)

type Repository interface {
	InsertUser(name, email, telephone, password, gender, age, img, role string) error
	GetByEmail(email string) (entity.User, error)
	GetLandingPageData() (LandingPageData, error)
}

type LandingPageData struct {
	PopularCategories []string
	BestPartner       []Partner
}

type Partner struct {
	Name     string `json:"nama"`
	Price    string `json:"harga"`
	Gender   string `json:"gender"`
	Category string `json:"category"`
	Rating   string `json:"rating"`
}

type TokenGenerator interface {
	GenerateToken(id, name, email, telephone, gender string) (string, error)
}

type ImageSaver interface {
	SaveImage(img io.ReadWriteCloser) (string, error)
}

type Mailer interface {
	Mail(to, title, body string) error
}

type Service struct {
	tokenGenerator TokenGenerator
	imageSaver     ImageSaver
	mailer         Mailer
	repo           Repository
}

func New(generator TokenGenerator, repo Repository, saver ImageSaver, mailer Mailer) *Service {
	return &Service{
		tokenGenerator: generator,
		repo:           repo,
		imageSaver:     saver,
		mailer:         mailer,
	}
}
