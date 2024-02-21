package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/models"
)

const Fieldwarehouse_home_redis = "LISTWAREHOUSE_BACKEND"
const Fieldwarehousestorage_home_redis = "LISTWAREHOUSE_STORAGE_BACKEND"

func Warehousehome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_warehouse)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	var obj entities.Model_warehouse
	var arraobj []entities.Model_warehouse
	var objbranch entities.Model_branchshare
	var arraobjbranch []entities.Model_branchshare
	render_page := time.Now()
	redisdata := ""
	if client.Branch_id != "" {
		redisdata = Fieldwarehouse_home_redis + "_" + strings.ToUpper(client.Branch_id)
	} else {
		redisdata = Fieldwarehouse_home_redis
	}
	resultredis, flag := helpers.GetRedis(redisdata)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listbranch_RD, _, _, _ := jsonparser.Get(jsonredis, "listbranch")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		warehouse_id, _ := jsonparser.GetString(value, "warehouse_id")
		warehouse_idbranch, _ := jsonparser.GetString(value, "warehouse_idbranch")
		warehouse_nmbranch, _ := jsonparser.GetString(value, "warehouse_nmbranch")
		warehouse_name, _ := jsonparser.GetString(value, "warehouse_name")
		warehouse_alamat, _ := jsonparser.GetString(value, "warehouse_alamat")
		warehouse_phone1, _ := jsonparser.GetString(value, "warehouse_phone1")
		warehouse_phone2, _ := jsonparser.GetString(value, "warehouse_phone2")
		warehouse_status, _ := jsonparser.GetString(value, "warehouse_status")
		warehouse_status_css, _ := jsonparser.GetString(value, "warehouse_status_css")
		warehouse_create, _ := jsonparser.GetString(value, "warehouse_create")
		warehouse_update, _ := jsonparser.GetString(value, "warehouse_update")

		obj.Warehouse_id = warehouse_id
		obj.Warehouse_idbranch = warehouse_idbranch
		obj.Warehouse_nmbranch = warehouse_nmbranch
		obj.Warehouse_name = warehouse_name
		obj.Warehouse_alamat = warehouse_alamat
		obj.Warehouse_phone1 = warehouse_phone1
		obj.Warehouse_phone2 = warehouse_phone2
		obj.Warehouse_status = warehouse_status
		obj.Warehouse_status_css = warehouse_status_css
		obj.Warehouse_create = warehouse_create
		obj.Warehouse_update = warehouse_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listbranch_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		branch_id, _ := jsonparser.GetString(value, "branch_id")
		branch_name, _ := jsonparser.GetString(value, "branch_name")

		objbranch.Branch_id = branch_id
		objbranch.Branch_name = branch_name
		arraobjbranch = append(arraobjbranch, objbranch)
	})
	if !flag {
		result, err := models.Fetch_warehouseHome(client.Branch_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(redisdata, result, 60*time.Minute)
		fmt.Println("WAREHOUSE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("WAREHOUSE CACHE")
		return c.JSON(fiber.Map{
			"status":     fiber.StatusOK,
			"message":    "Success",
			"record":     arraobj,
			"listbranch": arraobjbranch,
			"time":       time.Since(render_page).String(),
		})
	}
}
func Warehousestoragehome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_warehousestorage)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	var obj entities.Model_warehousestorage
	var arraobj []entities.Model_warehousestorage
	render_page := time.Now()
	redisdata := Fieldwarehouse_home_redis + "_" + strings.ToUpper(client.Warehouse_id)
	resultredis, flag := helpers.GetRedis(redisdata)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		warehousestorage_id, _ := jsonparser.GetString(value, "warehousestorage_id")
		warehousestorage_name, _ := jsonparser.GetString(value, "warehousestorage_name")
		warehousestorage_totalbin, _ := jsonparser.GetInt(value, "warehousestorage_totalbin")
		warehousestorage_status, _ := jsonparser.GetString(value, "warehousestorage_status")
		warehousestorage_status_css, _ := jsonparser.GetString(value, "warehousestorage_status_css")
		warehousestorage_create, _ := jsonparser.GetString(value, "warehousestorage_create")
		warehousestorage_update, _ := jsonparser.GetString(value, "warehousestorage_update")

		obj.Warehousestorage_id = warehousestorage_id
		obj.Warehousestorage_name = warehousestorage_name
		obj.Warehousestorage_totalbin = int(warehousestorage_totalbin)
		obj.Warehousestorage_status = warehousestorage_status
		obj.Warehousestorage_status_css = warehousestorage_status_css
		obj.Warehousestorage_create = warehousestorage_create
		obj.Warehousestorage_update = warehousestorage_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_warehouseStorage(client.Warehouse_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(redisdata, result, 60*time.Minute)
		fmt.Println("WAREHOUSE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("WAREHOUSE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func WarehousestorageBinhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_warehousestoragebin)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	var obj entities.Model_warehousestoragebin
	var arraobj []entities.Model_warehousestoragebin
	render_page := time.Now()
	redisdata := Fieldwarehousestorage_home_redis + "_" + strings.ToUpper(client.Storage_id)
	resultredis, flag := helpers.GetRedis(redisdata)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		warehousestoragebin_id, _ := jsonparser.GetInt(value, "warehousestoragebin_id")
		warehousestoragebin_iduom, _ := jsonparser.GetString(value, "warehousestoragebin_iduom")
		warehousestoragebin_name, _ := jsonparser.GetString(value, "warehousestoragebin_name")
		warehousestoragebin_totalcapacity, _ := jsonparser.GetFloat(value, "warehousestoragebin_totalcapacity")
		warehousestoragebin_maxcapacity, _ := jsonparser.GetFloat(value, "warehousestoragebin_maxcapacity")
		warehousestoragebin_status, _ := jsonparser.GetString(value, "warehousestoragebin_status")
		warehousestoragebin_status_css, _ := jsonparser.GetString(value, "warehousestoragebin_status_css")
		warehousestoragebin_create, _ := jsonparser.GetString(value, "warehousestoragebin_create")
		warehousestoragebin_update, _ := jsonparser.GetString(value, "warehousestoragebin_update")

		obj.Warehousestoragebin_id = int(warehousestoragebin_id)
		obj.Warehousestoragebin_iduom = warehousestoragebin_iduom
		obj.Warehousestoragebin_name = warehousestoragebin_name
		obj.Warehousestoragebin_totalcapacity = float32(warehousestoragebin_totalcapacity)
		obj.Warehousestoragebin_maxcapacity = float32(warehousestoragebin_maxcapacity)
		obj.Warehousestoragebin_status = warehousestoragebin_status
		obj.Warehousestoragebin_status_css = warehousestoragebin_status_css
		obj.Warehousestoragebin_create = warehousestoragebin_create
		obj.Warehousestoragebin_update = warehousestoragebin_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_warehouseStorageBin(client.Storage_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(redisdata, result, 60*time.Minute)
		fmt.Println("WAREHOUSE STORAGE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("WAREHOUSE STORAGE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}

func WarehouseSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_warehousesave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	// admin, idrecord, idbranch, name, alamat, phone1, phone2, status, sData string
	result, err := models.Save_warehouse(
		client_admin,
		client.Warehouse_id, client.Warehouse_idbranch, client.Warehouse_name, client.Warehouse_alamat,
		client.Warehouse_phone1, client.Warehouse_phone2, client.Warehouse_status,
		client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_warehouse(client.Warehouse_idbranch, "", "")
	return c.JSON(result)
}
func WarehouseStorageSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_warehousestoragesave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	// admin, idrecord, idwarehouse, name, status, sData string
	result, err := models.Save_warehousestorage(
		client_admin,
		client.Warehousestorage_id, client.Warehousestorage_idwarehouse, client.Warehousestorage_name, client.Warehousestorage_status,
		client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_warehouse("", client.Warehousestorage_idwarehouse, "")
	return c.JSON(result)
}
func WarehouseStorageBinSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_storagebinsave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	// admin, idstorage, iduom, name, status, sData string, idrecord int, maxcapacity float32
	result, err := models.Save_warehousestoragebin(
		client_admin,
		client.Storagebin_idstorage, client.Storagebin_iduom,
		client.Storagebin_name, client.Storagebin_status,
		client.Sdata, client.Storagebin_id, client.Storagebin_maxcapacity)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_warehouse("", "", client.Storagebin_idstorage)
	return c.JSON(result)
}
func _deleteredis_warehouse(idbranch, idwarehouse, idstorage string) {
	val_master_default := helpers.DeleteRedis(Fieldwarehouse_home_redis)
	fmt.Printf("Redis Delete BACKEND WAREHOUSE : %d\n", val_master_default)

	val_master := helpers.DeleteRedis(Fieldwarehouse_home_redis + "_" + strings.ToUpper(idbranch))
	fmt.Printf("Redis Delete BACKEND WAREHOUSE : %d\n", val_master)

	val_master_warehouse := helpers.DeleteRedis(Fieldwarehouse_home_redis + "_" + strings.ToUpper(idwarehouse))
	fmt.Printf("Redis Delete BACKEND WAREHOUSE : %d\n", val_master_warehouse)

	val_master_storage := helpers.DeleteRedis(Fieldwarehousestorage_home_redis + "_" + strings.ToUpper(idstorage))
	fmt.Printf("Redis Delete BACKEND WAREHOUSE : %d\n", val_master_storage)

}
