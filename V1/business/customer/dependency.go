package customer

type Repository interface {
	SeeAllPartners(withOptions map[string]string) ([]Partner, error)
	GetPartnerDetail(id string) (Partner, error)
}

type Partner struct {
	Name     string `json:"nama"`
	Price    string `json:"harga"`
	Gender   string `json:"gender"`
	Category string `json:"category"`
	Rating   string `json:"rating"`
}

type Service struct {
	repo Repository
}
