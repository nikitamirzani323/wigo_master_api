package entities

type Model_rfq struct {
	Rfq_id         string  `json:"rfq_id"`
	Rfq_date       string  `json:"rfq_date"`
	Rfq_idbranch   string  `json:"rfq_idbranch"`
	Rfq_idvendor   string  `json:"rfq_idvendor"`
	Rfq_idcurr     string  `json:"rfq_idcurr"`
	Rfq_tipedoc    string  `json:"rfq_tipedoc"`
	Rfq_nmbranch   string  `json:"rfq_nmbranch"`
	Rfq_nmvendor   string  `json:"rfq_nmvendor"`
	Rfq_totalitem  float64 `json:"rfq_totalitem"`
	Rfq_totalrfq   float64 `json:"rfq_totalrfq"`
	Rfq_status     string  `json:"rfq_status"`
	Rfq_status_css string  `json:"rfq_status_css"`
	Rfq_create     string  `json:"rfq_create"`
	Rfq_update     string  `json:"rfq_update"`
}
type Model_rfqdetail struct {
	Rfqdetail_id                      string  `json:"rfqdetail_id"`
	Rfqdetail_idpurchaserequestdetail string  `json:"rfqdetail_idpurchaserequestdetail"`
	Rfqdetail_idpurchaserequest       string  `json:"rfqdetail_idpurchaserequest"`
	Rfqdetail_nmdepartement           string  `json:"rfqdetail_nmdepartement"`
	Rfqdetail_nmemployee              string  `json:"rfqdetail_nmemployee"`
	Rfqdetail_iditem                  string  `json:"rfqdetail_iditem"`
	Rfqdetail_nmitem                  string  `json:"rfqdetail_nmitem"`
	Rfqdetail_descitem                string  `json:"rfqdetail_descitem"`
	Rfqdetail_qty                     float64 `json:"rfqdetail_qty"`
	Rfqdetail_iduom                   string  `json:"rfqdetail_iduom"`
	Rfqdetail_price                   float64 `json:"rfqdetail_price"`
	Rfqdetail_status                  string  `json:"rfqdetail_status"`
	Rfqdetail_status_css              string  `json:"rfqdetail_status_css"`
	Rfqdetail_create                  string  `json:"rfqdetail_create"`
	Rfqdetail_update                  string  `json:"rfqdetail_update"`
}

type Controller_rfqsave struct {
	Page           string  `json:"page" validate:"required"`
	Sdata          string  `json:"sdata" validate:"required"`
	Rfq_search     string  `json:"rfq_search"`
	Rfq_page       int     `json:"rfq_page"`
	Rfq_id         string  `json:"rfq_id"`
	Rfq_idbranch   string  `json:"rfq_idbranch" validate:"required"`
	Rfq_idvendor   string  `json:"rfq_idvendor" validate:"required"`
	Rfq_idcurr     string  `json:"rfq_idcurr" validate:"required"`
	Rfq_tipedoc    string  `json:"rfq_tipedoc" validate:"required"`
	Rfq_listdetail string  `json:"rfq_listdetail" validate:"required"`
	Rfq_totalitem  float32 `json:"rfq_totalitem" validate:"required"`
	Rfq_subtotal   float32 `json:"rfq_subtotal" validate:"required"`
}
type Controller_rfqdetailsave struct {
	Page                              string  `json:"page" validate:"required"`
	Sdata                             string  `json:"sdata" validate:"required"`
	Rfqdetail_search                  string  `json:"rfqdetail_search"`
	Rfqdetail_page                    int     `json:"rfqdetail_page"`
	Rfqdetail_id                      string  `json:"rfqdetail_id"`
	Rfqdetail_idpurchaserequestdetail string  `json:"rfqdetail_idpurchaserequestdetail" validate:"required"`
	Rfqdetail_idpurchaserequest       string  `json:"rfqdetail_idpurchaserequest" validate:"required"`
	Rfqdetail_iditem                  string  `json:"rfqdetail_iditem" validate:"required"`
	Rfqdetail_nmitem                  string  `json:"rfqdetail_nmitem" validate:"required"`
	Rfqdetail_descitem                string  `json:"rfqdetail_descitem"`
	Rfqdetail_qty                     float32 `json:"rfqdetail_qty" validate:"required"`
	Rfqdetail_iduom                   string  `json:"rfqdetail_iduom" validate:"required"`
	Rfqdetail_price                   float32 `json:"rfqdetail_price" validate:"required"`
}
type Controller_rfq struct {
	Rfq_search string `json:"rfq_search"`
	Rfq_page   int    `json:"rfq_page"`
}
type Controller_rfqdetail struct {
	Rfq_id string `json:"rfq_id" validate:"required"`
}
type Controller_rfqstatus struct {
	Rfq_id     string `json:"rfq_id" validate:"required"`
	Rfq_status string `json:"rfq_status" validate:"required"`
}
