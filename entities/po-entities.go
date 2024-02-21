package entities

type Model_po struct {
	Po_id         string  `json:"po_id"`
	Po_date       string  `json:"po_date"`
	Po_idrfq      string  `json:"po_idrfq"`
	Po_idbranch   string  `json:"po_idbranch"`
	Po_idvendor   string  `json:"po_idvendor"`
	Po_idcurr     string  `json:"po_idcurr"`
	Po_tipedoc    string  `json:"po_tipedoc"`
	Po_nmbranch   string  `json:"po_nmbranch"`
	Po_nmvendor   string  `json:"po_nmvendor"`
	Po_discount   float64 `json:"po_discount"`
	Po_ppn        float64 `json:"po_ppn"`
	Po_pph        float64 `json:"po_pph"`
	Po_totalitem  float64 `json:"po_totalitem"`
	Po_subtotal   float64 `json:"po_subtotal"`
	Po_grandtotal float64 `json:"po_grandtotal"`
	Po_status     string  `json:"po_status"`
	Po_status_css string  `json:"po_status_css"`
	Po_create     string  `json:"po_create"`
	Po_update     string  `json:"po_update"`
}
type Model_podetail struct {
	Podetail_id                      string  `json:"podetail_id"`
	Podetail_idpurchaserequestdetail string  `json:"podetail_idpurchaserequestdetail"`
	Podetail_idpurchaserequest       string  `json:"podetail_idpurchaserequest"`
	Podetail_nmdepartement           string  `json:"podetail_nmdepartement"`
	Podetail_nmemployee              string  `json:"podetail_nmemployee"`
	Podetail_iditem                  string  `json:"podetail_iditem"`
	Podetail_nmitem                  string  `json:"podetail_nmitem"`
	Podetail_descitem                string  `json:"podetail_descitem"`
	Podetail_qty                     float64 `json:"podetail_qty"`
	Podetail_iduom                   string  `json:"podetail_iduom"`
	Podetail_price                   float64 `json:"podetail_price"`
	Podetail_status                  string  `json:"podetail_status"`
	Podetail_status_css              string  `json:"podetail_status_css"`
	Podetail_create                  string  `json:"podetail_create"`
	Podetail_update                  string  `json:"podetail_update"`
}

type Controller_posave struct {
	Page          string  `json:"page" validate:"required"`
	Sdata         string  `json:"sdata" validate:"required"`
	Po_search     string  `json:"po_search"`
	Po_page       int     `json:"po_page"`
	Po_id         string  `json:"po_id"`
	Po_idrfq      string  `json:"po_idrfq" validate:"required"`
	Po_discount   float64 `json:"po_discount"`
	Po_ppn        float64 `json:"po_ppn" `
	Po_pph        float64 `json:"po_pph" `
	Po_ppn_total  float64 `json:"po_ppn_total" `
	Po_pph_total  float64 `json:"po_pph_total" `
	Po_totalitem  float64 `json:"po_totalitem" validate:"required"`
	Po_subtotal   float64 `json:"po_subtotal" validate:"required"`
	Po_grandtotal float64 `json:"po_grandtotal" validate:"required"`
}
type Controller_podetailsave struct {
	Page                             string  `json:"page" validate:"required"`
	Sdata                            string  `json:"sdata" validate:"required"`
	Podetail_search                  string  `json:"Podetail_search"`
	Podetail_page                    int     `json:"Podetail_page"`
	Podetail_id                      string  `json:"Podetail_id"`
	Podetail_idpurchaserequestdetail string  `json:"Podetail_idpurchaserequestdetail" validate:"required"`
	Podetail_idpurchaserequest       string  `json:"Podetail_idpurchaserequest" validate:"required"`
	Podetail_iditem                  string  `json:"Podetail_iditem" validate:"required"`
	Podetail_nmitem                  string  `json:"Podetail_nmitem" validate:"required"`
	Podetail_descitem                string  `json:"Podetail_descitem"`
	Podetail_qty                     float32 `json:"Podetail_qty" validate:"required"`
	Podetail_iduom                   string  `json:"Podetail_iduom" validate:"required"`
	Podetail_price                   float32 `json:"Podetail_price" validate:"required"`
}
type Controller_po struct {
	Po_search string `json:"po_search"`
	Po_page   int    `json:"po_page"`
}
type Controller_podetail struct {
	Po_id string `json:"po_id" validate:"required"`
}
type Controller_postatus struct {
	Po_id     string `json:"po_id" validate:"required"`
	Po_status string `json:"po_status" validate:"required"`
}
