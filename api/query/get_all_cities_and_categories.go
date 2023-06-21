package query

type Category struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type City struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type AllCitiesAndCategories struct {
	Cities     []City     `json:"cities"`
	Categories []Category `json:"categories"`
}
