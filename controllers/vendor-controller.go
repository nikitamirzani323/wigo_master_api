package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/WIGO_MASTER_API/entities"
	"github.com/nikitamirzani323/WIGO_MASTER_API/helpers"
	"github.com/nikitamirzani323/WIGO_MASTER_API/models"
)

const Fieldcatevendor_home_redis = "LISTCATEVENDOR_BACKEND"
const Fieldvendor_home_redis = "LISTVENDOR_BACKEND"

func Catevendorhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_catevendor)
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
	fmt.Println(client.Catevendor_page)
	if client.Catevendor_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldcatevendor_home_redis + "_" + strconv.Itoa(client.Catevendor_page) + "_" + client.Catevendor_search)
		fmt.Printf("Redis Delete BACKEND CATEVENDOR : %d", val_pattern)
	}

	var obj entities.Model_catevendor
	var arraobj []entities.Model_catevendor
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcatevendor_home_redis + "_" + strconv.Itoa(client.Catevendor_page) + "_" + client.Catevendor_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		catevendor_id, _ := jsonparser.GetInt(value, "catevendor_id")
		catevendor_name, _ := jsonparser.GetString(value, "catevendor_name")
		catevendor_status, _ := jsonparser.GetString(value, "catevendor_status")
		catevendor_status_css, _ := jsonparser.GetString(value, "catevendor_status_css")
		catevendor_create, _ := jsonparser.GetString(value, "catevendor_create")
		catevendor_update, _ := jsonparser.GetString(value, "catevendor_update")

		obj.Catevendor_id = int(catevendor_id)
		obj.Catevendor_name = catevendor_name
		obj.Catevendor_status = catevendor_status
		obj.Catevendor_status_css = catevendor_status_css
		obj.Catevendor_create = catevendor_create
		obj.Catevendor_update = catevendor_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_catevendorHome(client.Catevendor_search, client.Catevendor_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcatevendor_home_redis+"_"+strconv.Itoa(client.Catevendor_page)+"_"+client.Catevendor_search, result, 60*time.Minute)
		fmt.Println("CATE VENDOR MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("CATE VENDOR CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"time":        time.Since(render_page).String(),
		})
	}
}
func Vendorhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_vendor)
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
	fmt.Println(client.Vendor_page)
	if client.Vendor_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldvendor_home_redis + "_" + strconv.Itoa(client.Vendor_page) + "_" + client.Vendor_search)
		fmt.Printf("Redis Delete BACKEND CATEITEM : %d", val_pattern)
	}

	var obj entities.Model_vendor
	var arraobj []entities.Model_vendor
	var objcatevendor entities.Model_catevendorshare
	var arraobjcatevendor []entities.Model_catevendorshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldvendor_home_redis + "_" + strconv.Itoa(client.Vendor_page) + "_" + client.Vendor_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listcatevendor_RD, _, _, _ := jsonparser.Get(jsonredis, "listcatevendor")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		vendor_id, _ := jsonparser.GetString(value, "vendor_id")
		vendor_name, _ := jsonparser.GetString(value, "vendor_name")
		vendor_pic, _ := jsonparser.GetString(value, "vendor_pic")
		vendor_alamat, _ := jsonparser.GetString(value, "vendor_alamat")
		vendor_email, _ := jsonparser.GetString(value, "vendor_email")
		vendor_phone1, _ := jsonparser.GetString(value, "vendor_phone1")
		vendor_phone2, _ := jsonparser.GetString(value, "vendor_phone2")
		vendor_status, _ := jsonparser.GetString(value, "vendor_status")
		vendor_status_css, _ := jsonparser.GetString(value, "vendor_status_css")
		vendor_create, _ := jsonparser.GetString(value, "vendor_create")
		vendor_update, _ := jsonparser.GetString(value, "vendor_update")

		obj.Vendor_id = vendor_id
		obj.Vendor_name = vendor_name
		obj.Vendor_pic = vendor_pic
		obj.Vendor_alamat = vendor_alamat
		obj.Vendor_email = vendor_email
		obj.Vendor_phone1 = vendor_phone1
		obj.Vendor_phone2 = vendor_phone2
		obj.Vendor_status = vendor_status
		obj.Vendor_status_css = vendor_status_css
		obj.Vendor_create = vendor_create
		obj.Vendor_update = vendor_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcatevendor_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		catevendor_id, _ := jsonparser.GetInt(value, "catevendor_id")
		catevendor_name, _ := jsonparser.GetString(value, "catevendor_name")

		objcatevendor.Catevendor_id = int(catevendor_id)
		objcatevendor.Catevendor_name = catevendor_name
		arraobjcatevendor = append(arraobjcatevendor, objcatevendor)
	})
	if !flag {
		result, err := models.Fetch_vendorHome(client.Vendor_search, client.Vendor_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldvendor_home_redis+"_"+strconv.Itoa(client.Vendor_page)+"_"+client.Vendor_search, result, 60*time.Minute)
		fmt.Println("VENDOR MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("VENDOR CACHE")
		return c.JSON(fiber.Map{
			"status":         fiber.StatusOK,
			"message":        "Success",
			"record":         arraobj,
			"listcatevendor": arraobjcatevendor,
			"perpage":        perpage_RD,
			"totalrecord":    totalrecord_RD,
			"time":           time.Since(render_page).String(),
		})
	}
}
func Vendorshare(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_vendor)
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
	fmt.Println(client.Vendor_page)
	if client.Vendor_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldvendor_home_redis + "_" + client.Vendor_search)
		fmt.Printf("Redis Delete BACKEND VENDOR : %d", val_pattern)
	}

	var obj entities.Model_vendorshare
	var arraobj []entities.Model_vendorshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldvendor_home_redis + "_" + client.Vendor_search)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		vendor_id, _ := jsonparser.GetString(value, "vendor_id")
		vendor_nmcatevendor, _ := jsonparser.GetString(value, "vendor_nmcatevendor")
		vendor_name, _ := jsonparser.GetString(value, "vendor_name")

		obj.Vendor_id = vendor_id
		obj.Vendor_name = vendor_name
		obj.Vendor_nmcatevendor = vendor_nmcatevendor
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_vendorShare(client.Vendor_search)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldvendor_home_redis+"_"+client.Vendor_search, result, 60*time.Minute)
		fmt.Println("VENDOR SHARE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("VENDOR SHARE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CatevendorSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_catevendorsave)
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

	// admin, name, status, sData string, idrecord int
	result, err := models.Save_catevendor(
		client_admin,
		client.Catevendor_name, client.Catevendor_status,
		client.Sdata, client.Catevendor_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_vendor(client.Catevendor_search, client.Catevendor_page)
	return c.JSON(result)
}
func VendorSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_vendorsave)
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

	// admin, idrecord, name, pic, alamat, email, phone1, phone2, status, sData string, idcatevendor int
	result, err := models.Save_vendor(
		client_admin,
		client.Vendor_id, client.Vendor_name, client.Vendor_pic,
		client.Vendor_alamat, client.Vendor_email, client.Vendor_phone1, client.Vendor_phone2, client.Vendor_status,
		client.Sdata, client.Vendor_idcatevendor)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_vendor(client.Vendor_search, client.Vendor_page)
	return c.JSON(result)
}
func _deleteredis_vendor(search string, page int) {
	val_master_catevendor := helpers.DeleteRedis(Fieldcatevendor_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND CATEVENDOR : %d\n", val_master_catevendor)

	val_master := helpers.DeleteRedis(Fieldvendor_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND VENDOR : %d\n", val_master)

}
