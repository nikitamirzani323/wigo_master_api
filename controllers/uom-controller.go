package controllers

import (
	"fmt"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/WIGO_MASTER_API/entities"
	"github.com/nikitamirzani323/WIGO_MASTER_API/helpers"
	"github.com/nikitamirzani323/WIGO_MASTER_API/models"
)

const Fieluom_home_redis = "LISTUOM_BACKEND"
const Fieluomshare_home_redis = "LISTUOMSHARE_BACKEND"

func Uomhome(c *fiber.Ctx) error {
	var obj entities.Model_uom
	var arraobj []entities.Model_uom
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieluom_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		uom_id, _ := jsonparser.GetString(value, "uom_id")
		uom_name, _ := jsonparser.GetString(value, "uom_name")
		uom_status, _ := jsonparser.GetString(value, "uom_status")
		uom_status_css, _ := jsonparser.GetString(value, "uom_status_css")
		uom_create, _ := jsonparser.GetString(value, "uom_create")
		uom_update, _ := jsonparser.GetString(value, "uom_update")

		obj.Uom_id = uom_id
		obj.Uom_name = uom_name
		obj.Uom_status = uom_status
		obj.Uom_status_css = uom_status_css
		obj.Uom_create = uom_create
		obj.Uom_update = uom_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_uomHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieluom_home_redis, result, 60*time.Minute)
		fmt.Println("UOM MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("UOM CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Uomshare(c *fiber.Ctx) error {
	var obj entities.Model_uomshare
	var arraobj []entities.Model_uomshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieluom_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		uom_id, _ := jsonparser.GetString(value, "uom_id")
		uom_name, _ := jsonparser.GetString(value, "uom_name")

		obj.Uom_id = uom_id
		obj.Uom_name = uom_name
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_uomShare()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieluom_home_redis, result, 60*time.Minute)
		fmt.Println("UOM MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("UOM CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func UomSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_uomsave)
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
	result, err := models.Save_uom(
		client_admin,
		client.Uom_id, client.Uom_name, client.Uom_status, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_uom()
	return c.JSON(result)
}
func _deleteredis_uom() {
	val_master := helpers.DeleteRedis(Fieluom_home_redis)
	fmt.Printf("Redis Delete BACKEND UOM : %d", val_master)

	val_master_share := helpers.DeleteRedis(Fieluomshare_home_redis)
	fmt.Printf("Redis Delete BACKEND UOM SHARE : %d", val_master_share)

}
