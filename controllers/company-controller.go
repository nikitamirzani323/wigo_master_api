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

const Fieldcompany_home_redis = "LISTCOMPANY_BACKEND"

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
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompany_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
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
		fmt.Println("COMPANY MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY CACHE")
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

	_deleteredis_company()
	return c.JSON(result)
}
func _deleteredis_company() {
	val_master := helpers.DeleteRedis(Fieldcompany_home_redis)
	fmt.Printf("Redis Delete BACKEND COMPANY : %d", val_master)

}
