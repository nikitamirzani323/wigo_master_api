package entities

type Model_uom struct {
	Uom_id         string `json:"uom_id"`
	Uom_name       string `json:"uom_name"`
	Uom_status     string `json:"uom_status"`
	Uom_status_css string `json:"uom_status_css"`
	Uom_create     string `json:"uom_create"`
	Uom_update     string `json:"uom_update"`
}
type Model_uomshare struct {
	Uom_id   string `json:"uom_id"`
	Uom_name string `json:"uom_name"`
}
type Controller_uomsave struct {
	Page       string `json:"page" validate:"required"`
	Sdata      string `json:"sdata" validate:"required"`
	Uom_id     string `json:"uom_id"`
	Uom_name   string `json:"uom_name" validate:"required"`
	Uom_status string `json:"uom_status" validate:"required"`
}
