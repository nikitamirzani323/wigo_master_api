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

const Fielddepartement_home_redis = "LISTDEPARTEMENT_BACKEND"

func Departementhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_departement)
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
	fmt.Println(client.Departement_page)
	if client.Departement_search != "" {
		val_pattern := helpers.DeleteRedis(Fielddepartement_home_redis + "_" + strconv.Itoa(client.Departement_page) + "_" + client.Departement_search)
		fmt.Printf("Redis Delete BACKEND DEPARTEMENT : %d", val_pattern)
	}

	var obj entities.Model_departement
	var arraobj []entities.Model_departement
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielddepartement_home_redis + "_" + strconv.Itoa(client.Departement_page) + "_" + client.Departement_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		departement_id, _ := jsonparser.GetString(value, "departement_id")
		departement_name, _ := jsonparser.GetString(value, "departement_name")
		departement_status, _ := jsonparser.GetString(value, "departement_status")
		departement_status_css, _ := jsonparser.GetString(value, "departement_status_css")
		departement_create, _ := jsonparser.GetString(value, "departement_create")
		departement_update, _ := jsonparser.GetString(value, "departement_update")

		obj.Departement_id = departement_id
		obj.Departement_name = departement_name
		obj.Departement_status = departement_status
		obj.Departement_status_css = departement_status_css
		obj.Departement_create = departement_create
		obj.Departement_update = departement_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_departementHome(client.Departement_search, client.Departement_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fielddepartement_home_redis+"_"+strconv.Itoa(client.Departement_page)+"_"+client.Departement_search, result, 60*time.Minute)
		fmt.Println("DEPARTEMENT MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("DEPARTEMENT CACHE")
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
func DepartementSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_departementsave)
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

	// admin, idrecord, name, status, sData string
	result, err := models.Save_departement(
		client_admin,
		client.Departement_id, client.Departement_name, client.Departement_status, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_department(client.Departement_search, client.Departement_page)
	return c.JSON(result)
}
func _deleteredis_department(search string, page int) {
	val_master := helpers.DeleteRedis(Fielddepartement_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND DEPARTEMENT : %d\n", val_master)

}
