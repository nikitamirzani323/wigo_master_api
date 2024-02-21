package entities

type Model_departement struct {
	Departement_id         string `json:"departement_id"`
	Departement_name       string `json:"departement_name"`
	Departement_status     string `json:"departement_status"`
	Departement_status_css string `json:"departement_status_css"`
	Departement_create     string `json:"departement_create"`
	Departement_update     string `json:"departement_update"`
}
type Model_departementshare struct {
	Departement_id   string `json:"departement_id"`
	Departement_name string `json:"departement_name"`
}
type Controller_departementsave struct {
	Page               string `json:"page" validate:"required"`
	Sdata              string `json:"sdata" validate:"required"`
	Departement_search string `json:"departement_search"`
	Departement_page   int    `json:"departement_page"`
	Departement_id     string `json:"departement_id"`
	Departement_name   string `json:"departement_name" validate:"required"`
	Departement_status string `json:"departement_status" validate:"required"`
}
type Controller_departement struct {
	Departement_search string `json:"departement_search"`
	Departement_page   int    `json:"departement_page"`
}
