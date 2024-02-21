package entities

type Model_purchaserequest struct {
	Purchaserequest_id            string  `json:"purchaserequest_id"`
	Purchaserequest_date          string  `json:"purchaserequest_date"`
	Purchaserequest_idbranch      string  `json:"purchaserequest_idbranch"`
	Purchaserequest_iddepartement string  `json:"purchaserequest_iddepartement"`
	Purchaserequest_idemployee    string  `json:"purchaserequest_idemployee"`
	Purchaserequest_idcurr        string  `json:"purchaserequest_idcurr"`
	Purchaserequest_tipedoc       string  `json:"purchaserequest_tipedoc"`
	Purchaserequest_periodedoc    string  `json:"purchaserequest_periodedoc"`
	Purchaserequest_nmbranch      string  `json:"purchaserequest_nmbranch"`
	Purchaserequest_nmdepartement string  `json:"purchaserequest_nmdepartement"`
	Purchaserequest_nmemployee    string  `json:"purchaserequest_nmemployee"`
	Purchaserequest_totalitem     float64 `json:"purchaserequest_totalitem"`
	Purchaserequest_totalpr       float64 `json:"purchaserequest_totalpr"`
	Purchaserequest_totalpo       float64 `json:"purchaserequest_totalpo"`
	Purchaserequest_remark        string  `json:"purchaserequest_remark"`
	Purchaserequest_docexpire     string  `json:"purchaserequest_docexpire"`
	Purchaserequest_status        string  `json:"purchaserequest_status"`
	Purchaserequest_status_css    string  `json:"purchaserequest_status_css"`
	Purchaserequest_create        string  `json:"purchaserequest_create"`
	Purchaserequest_update        string  `json:"purchaserequest_update"`
}
type Model_purchaserequestdetail struct {
	Purchaserequestdetail_id                string  `json:"purchaserequestdetail_id"`
	Purchaserequestdetail_idpurchaserequest string  `json:"purchaserequestdetail_idpurchaserequest"`
	Purchaserequestdetail_iditem            string  `json:"purchaserequestdetail_iditem"`
	Purchaserequestdetail_nmitem            string  `json:"purchaserequestdetail_nmitem"`
	Purchaserequestdetail_descitem          string  `json:"purchaserequestdetail_descitem"`
	Purchaserequestdetail_purpose           string  `json:"purchaserequestdetail_purpose"`
	Purchaserequestdetail_qty               float32 `json:"purchaserequestdetail_qty"`
	Purchaserequestdetail_iduom             string  `json:"purchaserequestdetail_iduom"`
	Purchaserequestdetail_price             float32 `json:"purchaserequestdetail_price"`
	Purchaserequestdetail_status            string  `json:"purchaserequestdetail_status"`
	Purchaserequestdetail_status_css        string  `json:"purchaserequestdetail_status_css"`
	Purchaserequestdetail_create            string  `json:"purchaserequestdetail_create"`
	Purchaserequestdetail_update            string  `json:"purchaserequestdetail_update"`
}
type Model_prdetail_view struct {
	Prdetailview_id                string  `json:"prdetailview_id"`
	Prdetailview_idpurchaserequest string  `json:"prdetailview_idpurchaserequest"`
	Prdetailview_date              string  `json:"prdetailview_date"`
	Prdetailview_tipedoc           string  `json:"prdetailview_tipedoc"`
	Prdetailview_nmbranch          string  `json:"prdetailview_nmbranch"`
	Prdetailview_nmdepartement     string  `json:"prdetailview_nmdepartement"`
	Prdetailview_nmemployee        string  `json:"prdetailview_nmemployee"`
	Prdetailview_idcurr            string  `json:"prdetailview_idcurr"`
	Prdetailview_iditem            string  `json:"prdetailview_iditem"`
	Prdetailview_nmitem            string  `json:"prdetailview_nmitem"`
	Prdetailview_descitem          string  `json:"prdetailview_descitem"`
	Prdetailview_purpose           string  `json:"prdetailview_purpose"`
	Prdetailview_qty               float32 `json:"prdetailview_qty"`
	Prdetailview_qty_po            float32 `json:"prdetailview_qty_po"`
	Prdetailview_iduom             string  `json:"prdetailview_iduom"`
	Prdetailview_price             float32 `json:"prdetailview_price"`
	Prdetailview_status            string  `json:"prdetailview_status"`
	Prdetailview_status_css        string  `json:"prdetailview_status_css"`
}

type Controller_purchaserequestsave struct {
	Page                          string  `json:"page" validate:"required"`
	Sdata                         string  `json:"sdata" validate:"required"`
	Purchaserequest_search        string  `json:"purchaserequest_search"`
	Purchaserequest_page          int     `json:"purchaserequest_page"`
	Purchaserequest_id            string  `json:"purchaserequest_id"`
	Purchaserequest_idbranch      string  `json:"purchaserequest_idbranch" validate:"required"`
	Purchaserequest_iddepartement string  `json:"purchaserequest_iddepartement" validate:"required"`
	Purchaserequest_idemployee    string  `json:"purchaserequest_idemployee" validate:"required"`
	Purchaserequest_idcurr        string  `json:"purchaserequest_idcurr" validate:"required"`
	Purchaserequest_tipedoc       string  `json:"purchaserequest_tipedoc" validate:"required"`
	Purchaserequest_listdetail    string  `json:"purchaserequest_listdetail" validate:"required"`
	Purchaserequest_totalitem     float32 `json:"purchaserequest_totalitem" validate:"required"`
	Purchaserequest_subtotal      float32 `json:"purchaserequest_subtotal" validate:"required"`
	Purchaserequest_remark        string  `json:"purchaserequest_remark"`
}
type Controller_purchaserequestdetailsave struct {
	Page                                    string  `json:"page" validate:"required"`
	Sdata                                   string  `json:"sdata" validate:"required"`
	Purchaserequestdetail_search            string  `json:"purchaserequestdetail_search"`
	Purchaserequestdetail_page              int     `json:"purchaserequestdetail_page"`
	Purchaserequestdetail_id                string  `json:"purchaserequestdetail_id"`
	Purchaserequestdetail_idpurchaserequest string  `json:"purchaserequestdetail_idpurchaserequest" validate:"required"`
	Purchaserequestdetail_iditem            string  `json:"purchaserequestdetail_iditem" validate:"required"`
	Purchaserequestdetail_nmitem            string  `json:"purchaserequestdetail_nmitem" validate:"required"`
	Purchaserequestdetail_descitem          string  `json:"purchaserequestdetail_descitem" validate:"required"`
	Purchaserequestdetail_purpose           string  `json:"purchaserequestdetail_purpose" validate:"required"`
	Purchaserequestdetail_iduom             string  `json:"purchaserequestdetail_iduom" validate:"required"`
	Purchaserequestdetail_qty               float32 `json:"purchaserequestdetail_qty" validate:"required"`
	Purchaserequestdetail_estimateprice     float32 `json:"purchaserequestdetail_estimateprice" validate:"required"`
}
type Controller_purchaserequest struct {
	Purchaserequest_search string `json:"purchaserequest_search"`
	Purchaserequest_status string `json:"purchaserequest_status"`
	Purchaserequest_page   int    `json:"purchaserequest_page"`
}
type Controller_purchaserequestdetail struct {
	Purchaserequest_id string `json:"purchaserequest_id" validate:"required"`
}
type Controller_prdetail_view struct {
	Purchaserequest_tipedoc  string `json:"purchaserequest_tipedoc" validate:"required"`
	Purchaserequest_idbranch string `json:"purchaserequest_idbranch" validate:"required"`
}
type Controller_purchaserequeststatus struct {
	Purchaserequest_id     string `json:"purchaserequest_id" validate:"required"`
	Purchaserequest_status string `json:"purchaserequest_status" validate:"required"`
}
