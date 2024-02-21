package entities

type Model_employee struct {
	Employee_id            string `json:"employee_id"`
	Employee_iddepartement string `json:"employee_iddepartement"`
	Employee_nmdepartement string `json:"employee_nmdepartement"`
	Employee_name          string `json:"employee_name"`
	Employee_alamat        string `json:"employee_alamat"`
	Employee_email         string `json:"employee_email"`
	Employee_phone1        string `json:"employee_phone1"`
	Employee_phone2        string `json:"employee_phone2"`
	Employee_status        string `json:"employee_status"`
	Employee_status_css    string `json:"employee_status_css"`
	Employee_create        string `json:"employee_create"`
	Employee_update        string `json:"employee_update"`
}
type Model_employeeshare struct {
	Employee_id   string `json:"employee_id"`
	Employee_name string `json:"employee_name"`
}
type Controller_employeesave struct {
	Page                   string `json:"page" validate:"required"`
	Sdata                  string `json:"sdata" validate:"required"`
	Employee_search        string `json:"employee_search"`
	Employee_page          int    `json:"employee_page"`
	Employee_id            string `json:"employee_id"`
	Employee_iddepartement string `json:"employee_iddepartement" validate:"required"`
	Employee_name          string `json:"employee_name" validate:"required"`
	Employee_alamat        string `json:"employee_alamat"`
	Employee_email         string `json:"employee_email"`
	Employee_phone1        string `json:"employee_phone1" validate:"required"`
	Employee_phone2        string `json:"employee_phone2"`
	Employee_status        string `json:"employee_status" validate:"required"`
}
type Controller_employee struct {
	Employee_search string `json:"employee_search"`
	Employee_page   int    `json:"employee_page"`
}
type Controller_employeeshare struct {
	Employee_iddepartement string `json:"employee_iddepartement"`
}
