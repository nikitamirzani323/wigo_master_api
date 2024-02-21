package entities

type Model_catevendor struct {
	Catevendor_id         int    `json:"catevendor_id"`
	Catevendor_name       string `json:"catevendor_name"`
	Catevendor_status     string `json:"catevendor_status"`
	Catevendor_status_css string `json:"catevendor_status_css"`
	Catevendor_create     string `json:"catevendor_create"`
	Catevendor_update     string `json:"catevendor_update"`
}
type Model_catevendorshare struct {
	Catevendor_id   int    `json:"catevendor_id"`
	Catevendor_name string `json:"catevendor_name"`
}
type Model_vendor struct {
	Vendor_id           string `json:"vendor_id"`
	Vendor_idcatevendor int    `json:"vendor_idcatevendor"`
	Vendor_nmcatevendor string `json:"vendor_nmcatevendor"`
	Vendor_name         string `json:"vendor_name"`
	Vendor_pic          string `json:"vendor_pic"`
	Vendor_alamat       string `json:"vendor_alamat"`
	Vendor_email        string `json:"vendor_email"`
	Vendor_phone1       string `json:"vendor_phone1"`
	Vendor_phone2       string `json:"vendor_phone2"`
	Vendor_status       string `json:"vendor_status"`
	Vendor_status_css   string `json:"vendor_status_css"`
	Vendor_create       string `json:"vendor_create"`
	Vendor_update       string `json:"vendor_update"`
}
type Model_vendorshare struct {
	Vendor_id           string `json:"vendor_id"`
	Vendor_nmcatevendor string `json:"vendor_nmcatevendor"`
	Vendor_name         string `json:"vendor_name"`
}

type Controller_catevendorsave struct {
	Page              string `json:"page" validate:"required"`
	Sdata             string `json:"sdata" validate:"required"`
	Catevendor_search string `json:"catevendor_search"`
	Catevendor_page   int    `json:"catevendor_page"`
	Catevendor_id     int    `json:"catevendor_id"`
	Catevendor_name   string `json:"catevendor_name" validate:"required"`
	Catevendor_status string `json:"catevendor_status" validate:"required"`
}
type Controller_vendorsave struct {
	Page                string `json:"page" validate:"required"`
	Sdata               string `json:"sdata" validate:"required"`
	Vendor_search       string `json:"vendor_search"`
	Vendor_page         int    `json:"vendor_page"`
	Vendor_id           string `json:"vendor_id"`
	Vendor_idcatevendor int    `json:"vendor_idcatevendor" validate:"required"`
	Vendor_name         string `json:"vendor_name" validate:"required"`
	Vendor_pic          string `json:"vendor_pic" validate:"required"`
	Vendor_alamat       string `json:"vendor_alamat"`
	Vendor_email        string `json:"vendor_email"`
	Vendor_phone1       string `json:"vendor_phone1" validate:"required"`
	Vendor_phone2       string `json:"vendor_phone2"`
	Vendor_status       string `json:"vendor_status" validate:"required"`
}
type Controller_catevendor struct {
	Catevendor_search string `json:"catevendor_search"`
	Catevendor_page   int    `json:"catevendor_page"`
}
type Controller_vendor struct {
	Vendor_search string `json:"vendor_search"`
	Vendor_page   int    `json:"vendor_page"`
}
