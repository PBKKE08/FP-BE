package repo

import (
	"database/sql"
	"errors"
	"github.com/PBKKE08/FP-BE/V1/business/customer"
	"github.com/rs/zerolog/log"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) SeeAllPartners(withOptions map[string]string) ([]customer.Partner, error){
	var partner customer.Partner
	var partners []customer.Partner 

	sqlStatement := `SELECT user.name, price, gender, category.name, rating FROM users RIGHT JOIN partners ON users.id = partners.user_id JOIN category ON category.id = partners.category_id`

	rows, err := r.db.Query(sqlStatement)

	if err != nil {
		return partners, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&partner.Name, &partner.Price, &partner.Gender, &partner.Category, &partner.Rating)
		if err != nil {
			return partners, err
		}

		partners = append(partners, partner)
	}

	return partners,nil
}

func (r *CustomerRepository) GetPartnerDetail(id string) (customer.Partner, error){
	q := `SELECT user.name, price, gender, category.name, rating FROM users JOIN partners ON users.id = partners.user_id JOIN category ON category.id = partners.category_id WHERE users.id = $1`
	var partner customer.Partner

	err := r.db.QueryRow(q, id).Scan(&partner.Name, &partner.Price, &partner.Gender, &partner.Category, &partner.Rating)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return partner, nil
		}

		log.Debug().Err(err).Stack()
		return partner, ErrUnknown
	}

	return partner, nil
}