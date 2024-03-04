package entities

type Model_company struct {
	Company_id         string  `json:"company_id"`
	Company_startjoin  string  `json:"company_startjoin"`
	Company_endjoin    string  `json:"company_endjoin"`
	Company_idcurr     string  `json:"company_idcurr"`
	Company_name       string  `json:"company_name"`
	Company_owner      string  `json:"company_owner"`
	Company_phone1     string  `json:"company_phone1"`
	Company_phone2     string  `json:"company_phone2"`
	Company_email      string  `json:"company_email"`
	Company_minfee     float64 `json:"company_minfee"`
	Company_url1       string  `json:"company_url1"`
	Company_url2       string  `json:"company_url2"`
	Company_status     string  `json:"company_status"`
	Company_status_css string  `json:"company_status_css"`
	Company_create     string  `json:"company_create"`
	Company_update     string  `json:"company_update"`
}
type Model_companyadmin struct {
	Companyadmin_id         int    `json:"companyadmin_id"`
	Companyadmin_idcompany  string `json:"companyadmin_idcompany"`
	Companyadmin_username   string `json:"companyadmin_username"`
	Companyadmin_name       string `json:"companyadmin_name"`
	Companyadmin_status     string `json:"companyadmin_status"`
	Companyadmin_status_css string `json:"companyadmin_status_css"`
	Companyadmin_create     string `json:"companyadmin_create"`
	Companyadmin_update     string `json:"companyadmin_update"`
}
type Model_companyshare struct {
	Company_id   string `json:"company_id"`
	Company_name string `json:"company_name"`
}

type Controller_companysave struct {
	Page           string  `json:"page" validate:"required"`
	Sdata          string  `json:"sdata" validate:"required"`
	Company_search string  `json:"company_search"`
	Company_page   int     `json:"company_page"`
	Company_id     string  `json:"company_id"`
	Company_idcurr string  `json:"company_idcurr" validate:"required"`
	Company_name   string  `json:"company_name" validate:"required"`
	Company_owner  string  `json:"company_owner" validate:"required"`
	Company_phone1 string  `json:"company_phone1"`
	Company_phone2 string  `json:"company_phone2"`
	Company_email  string  `json:"company_email"`
	Company_minfee float64 `json:"company_minfee"`
	Company_url1   string  `json:"company_url1" validate:"required"`
	Company_url2   string  `json:"company_url2" validate:"required"`
	Company_status string  `json:"company_status" validate:"required"`
}
type Controller_companyadminsave struct {
	Page                   string `json:"page" validate:"required"`
	Sdata                  string `json:"sdata" validate:"required"`
	Companyadmin_id        int    `json:"companyadmin_id"`
	Companyadmin_idcompany string `json:"companyadmin_idcompany" validate:"required"`
	Companyadmin_username  string `json:"companyadmin_username" validate:"required"`
	Companyadmin_password  string `json:"companyadmin_password`
	Companyadmin_name      string `json:"companyadmin_name" validate:"required"`
	Companyadmin_status    string `json:"companyadmin_status" validate:"required"`
}
type Controller_company struct {
	Company_search string `json:"company_search"`
	Company_page   int    `json:"company_page"`
}
type Controller_companyadmin struct {
	Companyadmin_idcompany string `json:"companyadmin_idcompany" `
}
