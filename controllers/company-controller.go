package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/WIGO_MASTER_API/entities"
	"github.com/nikitamirzani323/WIGO_MASTER_API/helpers"
	"github.com/nikitamirzani323/WIGO_MASTER_API/models"
)

const Fieldcompany_home_redis = "LISTCOMPANY_BACKEND"
const Fieldcompanyadmin_home_redis = "LISTCOMPANYADMIN_BACKEND"

func Companyhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_company)
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
	fmt.Println(client.Company_page)
	if client.Company_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldcompany_home_redis + "_" + strconv.Itoa(client.Company_page) + "_" + client.Company_search)
		fmt.Printf("Redis Delete BACKEND COMPANY : %d", val_pattern)
	}

	var obj entities.Model_company
	var arraobj []entities.Model_company
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompany_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listcurr_RD, _, _, _ := jsonparser.Get(jsonredis, "listcurr")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_id, _ := jsonparser.GetString(value, "company_id")
		company_startjoin, _ := jsonparser.GetString(value, "company_startjoin")
		company_endjoin, _ := jsonparser.GetString(value, "company_endjoin")
		company_idcurr, _ := jsonparser.GetString(value, "company_idcurr")
		company_name, _ := jsonparser.GetString(value, "company_name")
		company_owner, _ := jsonparser.GetString(value, "company_owner")
		company_phone1, _ := jsonparser.GetString(value, "company_phone1")
		company_phone2, _ := jsonparser.GetString(value, "company_phone2")
		company_email, _ := jsonparser.GetString(value, "company_email")
		company_minfee, _ := jsonparser.GetFloat(value, "company_minfee")
		company_url1, _ := jsonparser.GetString(value, "company_url1")
		company_url2, _ := jsonparser.GetString(value, "company_url2")
		company_status, _ := jsonparser.GetString(value, "company_status")
		company_status_css, _ := jsonparser.GetString(value, "company_status_css")
		company_create, _ := jsonparser.GetString(value, "company_create")
		company_update, _ := jsonparser.GetString(value, "company_update")

		obj.Company_id = company_id
		obj.Company_startjoin = company_startjoin
		obj.Company_endjoin = company_endjoin
		obj.Company_idcurr = company_idcurr
		obj.Company_name = company_name
		obj.Company_owner = company_owner
		obj.Company_email = company_email
		obj.Company_phone1 = company_phone1
		obj.Company_phone2 = company_phone2
		obj.Company_minfee = float64(company_minfee)
		obj.Company_url1 = company_url1
		obj.Company_url2 = company_url2
		obj.Company_status = company_status
		obj.Company_status_css = company_status_css
		obj.Company_create = company_create
		obj.Company_update = company_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcurr_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		curr_id, _ := jsonparser.GetString(value, "curr_id")

		objcurr.Curr_id = curr_id
		arraobjcurr = append(arraobjcurr, objcurr)
	})
	if !flag {
		result, err := models.Fetch_companyHome(client.Company_search, client.Company_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompany_home_redis+"_"+strconv.Itoa(client.Company_page)+"_"+client.Company_search, result, 60*time.Minute)
		fmt.Println("COMPANY DATABASE")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY CACHE")
		return c.JSON(fiber.Map{
			"status":   fiber.StatusOK,
			"message":  "Success",
			"record":   arraobj,
			"listcurr": arraobjcurr,
			"time":     time.Since(render_page).String(),
		})
	}
}
func Companyadminhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyadmin)
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

	var obj entities.Model_companyadmin
	var arraobj []entities.Model_companyadmin
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompanyadmin_home_redis + "_" + strings.ToLower(client.Companyadmin_idcompany))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyadmin_id, _ := jsonparser.GetInt(value, "companyadmin_id")
		companyadmin_idcompany, _ := jsonparser.GetString(value, "companyadmin_idcompany")
		companyadmin_username, _ := jsonparser.GetString(value, "companyadmin_username")
		companyadmin_name, _ := jsonparser.GetString(value, "companyadmin_name")
		companyadmin_status, _ := jsonparser.GetString(value, "companyadmin_status")
		companyadmin_status_css, _ := jsonparser.GetString(value, "companyadmin_status_css")
		companyadmin_create, _ := jsonparser.GetString(value, "companyadmin_create")
		companyadmin_update, _ := jsonparser.GetString(value, "companyadmin_update")

		obj.Companyadmin_id = int(companyadmin_id)
		obj.Companyadmin_idcompany = companyadmin_idcompany
		obj.Companyadmin_username = companyadmin_username
		obj.Companyadmin_name = companyadmin_name
		obj.Companyadmin_status = companyadmin_status
		obj.Companyadmin_status_css = companyadmin_status_css
		obj.Companyadmin_create = companyadmin_create
		obj.Companyadmin_update = companyadmin_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_companyadminHome(client.Companyadmin_idcompany)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanyadmin_home_redis+"_"+strings.ToLower(client.Companyadmin_idcompany), result, 60*time.Minute)
		fmt.Println("COMPANY ADMIN DATABASE")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY ADMIN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanySave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companysave)
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

	// admin, idrecord, idcurr, nmcompany, nmowner,
	// emailowner, phone1, phone2, url1, url2, status, sData string, minfee float64
	result, err := models.Save_company(
		client_admin,
		client.Company_id, client.Company_idcurr, client.Company_name, client.Company_owner,
		client.Company_email, client.Company_phone1, client.Company_phone2,
		client.Company_url1, client.Company_url2, client.Company_status,
		client.Sdata, client.Company_minfee)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company("")
	return c.JSON(result)
}
func CompanyadminSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyadminsave)
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

	// admin, idcompany, username, password, name, status, sData string, idrecord int
	result, err := models.Save_companyadmin(
		client_admin,
		client.Companyadmin_idcompany, client.Companyadmin_username, client.Companyadmin_password,
		client.Companyadmin_name, client.Companyadmin_status, client.Sdata, client.Companyadmin_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company(client.Companyadmin_idcompany)
	return c.JSON(result)
}
func _deleteredis_company(idcompany string) {
	val_master := helpers.DeleteRedis(Fieldcompany_home_redis)
	fmt.Printf("Redis Delete BACKEND COMPANY : %d", val_master)

	val_compadmin := helpers.DeleteRedis(Fieldcompanyadmin_home_redis + "_" + strings.ToLower(idcompany))
	fmt.Printf("Redis Delete BACKEND COMPANY ADMIN : %d", val_compadmin)

}
