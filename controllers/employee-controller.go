package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/models"
)

const Fieldemployee_home_redis = "LISTEMPLOYEE_BACKEND"
const Fieldemployeeshare_home_redis = "LISTEMPLOYEESHARE_BACKEND"

func Employeehome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_employee)
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
	fmt.Println(client.Employee_page)
	if client.Employee_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldemployee_home_redis + "_" + strconv.Itoa(client.Employee_page) + "_" + client.Employee_search)
		fmt.Printf("Redis Delete BACKEND DEPARTEMENT : %d", val_pattern)
	}

	var obj entities.Model_employee
	var arraobj []entities.Model_employee
	var objdepartement entities.Model_departementshare
	var arraobjdepartement []entities.Model_departementshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldemployee_home_redis + "_" + strconv.Itoa(client.Employee_page) + "_" + client.Employee_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listdepartement_RD, _, _, _ := jsonparser.Get(jsonredis, "listdepartement")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		employee_id, _ := jsonparser.GetString(value, "employee_id")
		employee_iddepartement, _ := jsonparser.GetString(value, "employee_iddepartement")
		employee_nmdepartement, _ := jsonparser.GetString(value, "employee_nmdepartement")
		employee_name, _ := jsonparser.GetString(value, "employee_name")
		employee_alamat, _ := jsonparser.GetString(value, "employee_alamat")
		employee_email, _ := jsonparser.GetString(value, "employee_email")
		employee_phone1, _ := jsonparser.GetString(value, "employee_phone1")
		employee_phone2, _ := jsonparser.GetString(value, "employee_phone2")
		employee_status, _ := jsonparser.GetString(value, "employee_status")
		employee_status_css, _ := jsonparser.GetString(value, "employee_status_css")
		employee_create, _ := jsonparser.GetString(value, "employee_create")
		employee_update, _ := jsonparser.GetString(value, "employee_update")

		obj.Employee_id = employee_id
		obj.Employee_iddepartement = employee_iddepartement
		obj.Employee_nmdepartement = employee_nmdepartement
		obj.Employee_name = employee_name
		obj.Employee_alamat = employee_alamat
		obj.Employee_email = employee_email
		obj.Employee_phone1 = employee_phone1
		obj.Employee_phone2 = employee_phone2
		obj.Employee_status = employee_status
		obj.Employee_status_css = employee_status_css
		obj.Employee_create = employee_create
		obj.Employee_update = employee_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listdepartement_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		departement_id, _ := jsonparser.GetString(value, "departement_id")
		departement_name, _ := jsonparser.GetString(value, "departement_name")

		objdepartement.Departement_id = departement_id
		objdepartement.Departement_name = departement_name
		arraobjdepartement = append(arraobjdepartement, objdepartement)
	})
	if !flag {
		result, err := models.Fetch_employeeHome(client.Employee_search, client.Employee_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldemployee_home_redis+"_"+strconv.Itoa(client.Employee_page)+"_"+client.Employee_search, result, 60*time.Minute)
		fmt.Println("EMPLOYEE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("EMPLOYEE CACHE")
		return c.JSON(fiber.Map{
			"status":          fiber.StatusOK,
			"message":         "Success",
			"record":          arraobj,
			"listdepartement": arraobjdepartement,
			"perpage":         perpage_RD,
			"totalrecord":     totalrecord_RD,
			"time":            time.Since(render_page).String(),
		})
	}
}
func Employeeshare(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_employeeshare)
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

	var obj entities.Model_employeeshare
	var arraobj []entities.Model_employeeshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldemployeeshare_home_redis + "_" + client.Employee_iddepartement)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		employee_id, _ := jsonparser.GetString(value, "employee_id")
		employee_name, _ := jsonparser.GetString(value, "employee_name")

		obj.Employee_id = employee_id
		obj.Employee_name = employee_name
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_employeeShare(client.Employee_iddepartement)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldemployeeshare_home_redis+"_"+client.Employee_iddepartement, result, 60*time.Minute)
		fmt.Println("EMPLOYEE SHARE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("EMPLOYEE SHARE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func EmployeeSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_employeesave)
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

	// admin, idrecord, iddepart, name, alamat, email, phone1, phone2, status, sData string
	result, err := models.Save_employee(
		client_admin,
		client.Employee_id, client.Employee_iddepartement, client.Employee_name,
		client.Employee_alamat, client.Employee_email, client.Employee_phone1, client.Employee_phone2,
		client.Employee_status, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_employee(client.Employee_search, client.Employee_page)
	return c.JSON(result)
}
func _deleteredis_employee(search string, page int) {
	val_master := helpers.DeleteRedis(Fieldemployee_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND EMPLOYEE : %d\n", val_master)

	val_master_share := helpers.DeleteRedis(Fieldemployeeshare_home_redis + "_" + search)
	fmt.Printf("Redis Delete BACKEND EMPLOYEE SHARE : %d\n", val_master_share)

}
