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

const Fieldbranch_home_redis = "LISTBRANCH_BACKEND"

func Branchhome(c *fiber.Ctx) error {
	var obj entities.Model_branch
	var arraobj []entities.Model_branch
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldbranch_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		branch_id, _ := jsonparser.GetString(value, "branch_id")
		branch_name, _ := jsonparser.GetString(value, "branch_name")
		branch_status, _ := jsonparser.GetString(value, "branch_status")
		branch_status_css, _ := jsonparser.GetString(value, "branch_status_css")
		branch_create, _ := jsonparser.GetString(value, "branch_create")
		branch_update, _ := jsonparser.GetString(value, "branch_update")

		obj.Branch_id = branch_id
		obj.Branch_name = branch_name
		obj.Branch_status = branch_status
		obj.Branch_status_css = branch_status_css
		obj.Branch_create = branch_create
		obj.Branch_update = branch_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_branchHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldbranch_home_redis, result, 60*time.Minute)
		fmt.Println("BRANCH MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("BRANCH CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func BranchSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_branchsave)
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

	// admin, idrecord, name, sData string
	result, err := models.Save_branch(
		client_admin,
		client.Branch_id, client.Branch_name, client.Branch_status, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_branch()
	return c.JSON(result)
}
func _deleteredis_branch() {
	val_master := helpers.DeleteRedis(Fieldbranch_home_redis)
	fmt.Printf("Redis Delete BACKEND BRANCH : %d", val_master)

}
