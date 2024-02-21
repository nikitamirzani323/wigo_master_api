package entities

type Model_warehouse struct {
	Warehouse_id         string `json:"warehouse_id"`
	Warehouse_idbranch   string `json:"warehouse_idbranch"`
	Warehouse_nmbranch   string `json:"warehouse_nmbranch"`
	Warehouse_name       string `json:"warehouse_name"`
	Warehouse_alamat     string `json:"warehouse_alamat"`
	Warehouse_phone1     string `json:"warehouse_phone1"`
	Warehouse_phone2     string `json:"warehouse_phone2"`
	Warehouse_status     string `json:"warehouse_status"`
	Warehouse_status_css string `json:"warehouse_status_css"`
	Warehouse_create     string `json:"warehouse_create"`
	Warehouse_update     string `json:"warehouse_update"`
}
type Model_warehousestorage struct {
	Warehousestorage_id         string `json:"warehousestorage_id"`
	Warehousestorage_name       string `json:"warehousestorage_name"`
	Warehousestorage_totalbin   int    `json:"warehousestorage_totalbin"`
	Warehousestorage_status     string `json:"warehousestorage_status"`
	Warehousestorage_status_css string `json:"warehousestorage_status_css"`
	Warehousestorage_create     string `json:"warehousestorage_create"`
	Warehousestorage_update     string `json:"warehousestorage_update"`
}
type Model_warehousestoragebin struct {
	Warehousestoragebin_id            int     `json:"warehousestoragebin_id"`
	Warehousestoragebin_iduom         string  `json:"warehousestoragebin_iduom"`
	Warehousestoragebin_name          string  `json:"warehousestoragebin_name"`
	Warehousestoragebin_totalcapacity float32 `json:"warehousestoragebin_totalcapacity"`
	Warehousestoragebin_maxcapacity   float32 `json:"warehousestoragebin_maxcapacity"`
	Warehousestoragebin_status        string  `json:"warehousestoragebin_status"`
	Warehousestoragebin_status_css    string  `json:"warehousestoragebin_status_css"`
	Warehousestoragebin_create        string  `json:"warehousestoragebin_create"`
	Warehousestoragebin_update        string  `json:"warehousestoragebin_update"`
}
type Controller_warehousesave struct {
	Page               string `json:"page" validate:"required"`
	Sdata              string `json:"sdata" validate:"required"`
	Warehouse_id       string `json:"warehouse_id" validate:"required"`
	Warehouse_idbranch string `json:"warehouse_idbranch" validate:"required"`
	Warehouse_name     string `json:"warehouse_name" validate:"required"`
	Warehouse_alamat   string `json:"warehouse_alamat" validate:"required"`
	Warehouse_phone1   string `json:"warehouse_phone1" `
	Warehouse_phone2   string `json:"warehouse_phone2" `
	Warehouse_status   string `json:"warehouse_status" validate:"required"`
}
type Controller_warehousestoragesave struct {
	Page                         string `json:"page" validate:"required"`
	Sdata                        string `json:"sdata" validate:"required"`
	Warehousestorage_id          string `json:"warehousestorage_id" validate:"required"`
	Warehousestorage_idwarehouse string `json:"warehousestorage_idwarehouse" validate:"required"`
	Warehousestorage_name        string `json:"warehousestorage_name" validate:"required"`
	Warehousestorage_status      string `json:"warehousestorage_status" validate:"required"`
}
type Controller_storagebinsave struct {
	Page                   string  `json:"page" validate:"required"`
	Sdata                  string  `json:"sdata" validate:"required"`
	Storagebin_id          int     `json:"storagebin_id" `
	Storagebin_idstorage   string  `json:"storagebin_idstorage" validate:"required"`
	Storagebin_iduom       string  `json:"storagebin_iduom" validate:"required"`
	Storagebin_name        string  `json:"storagebin_name" validate:"required"`
	Storagebin_maxcapacity float32 `json:"storagebin_maxcapacity" `
	Storagebin_status      string  `json:"storagebin_status" validate:"required"`
}
type Controller_warehouse struct {
	Branch_id string `json:"branch_id" `
}
type Controller_warehousestorage struct {
	Warehouse_id string `json:"warehouse_id" validate:"required"`
}
type Controller_warehousestoragebin struct {
	Storage_id string `json:"storage_id" validate:"required"`
}
