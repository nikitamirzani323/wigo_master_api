package entities

type Model_domain struct {
	Domain_id         int    `json:"domain_id"`
	Domain_name       string `json:"domain_name"`
	Domain_tipe       string `json:"domain_tipe"`
	Domain_status     string `json:"domain_status"`
	Domain_status_css string `json:"domain_status_css"`
	Domain_create     string `json:"domain_create"`
	Domain_update     string `json:"domain_update"`
}
type Model_checkdomain struct {
	Domain_name   string `json:"domain_name"`
	Domain_tipe   string `json:"domain_tipe"`
	Domain_status string `json:"domain_status"`
}
type Controller_domainsave struct {
	Page          string `json:"page" validate:"required"`
	Sdata         string `json:"sdata" validate:"required"`
	Domain_id     int    `json:"domain_id"`
	Domain_name   string `json:"domain_name" validate:"required"`
	Domain_tipe   string `json:"domain_tipe" validate:"required"`
	Domain_status string `json:"domain_status" validate:"required"`
}
type Controller_domaincheck struct {
	Domain string `json:"domain" validate:"required"`
	Tipe   string `json:"tipe" validate:"required"`
}
