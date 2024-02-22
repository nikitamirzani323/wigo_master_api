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

const Fieldrfq_home_redis = "RFQ_BACKEND"

func Rfqhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_rfq)
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

	if client.Rfq_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldrfq_home_redis + "_" + strconv.Itoa(client.Rfq_page) + "_" + client.Rfq_search)
		fmt.Printf("Redis Delete BACKEND RFQ : %d", val_pattern)
	}

	var obj entities.Model_rfq
	var arraobj []entities.Model_rfq
	var objbranch entities.Model_branchshare
	var arraobjbranch []entities.Model_branchshare
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldrfq_home_redis + "_" + strconv.Itoa(client.Rfq_page) + "_" + client.Rfq_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	Listbranch_RD, _, _, _ := jsonparser.Get(jsonredis, "listbranch")
	Listcurr_RD, _, _, _ := jsonparser.Get(jsonredis, "listcurr")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		rfq_id, _ := jsonparser.GetString(value, "rfq_id")
		rfq_date, _ := jsonparser.GetString(value, "rfq_date")
		rfq_idbranch, _ := jsonparser.GetString(value, "rfq_idbranch")
		rfq_idvendor, _ := jsonparser.GetString(value, "rfq_idvendor")
		rfq_idcurr, _ := jsonparser.GetString(value, "rfq_idcurr")
		rfq_tipedoc, _ := jsonparser.GetString(value, "rfq_tipedoc")
		rfq_nmbranch, _ := jsonparser.GetString(value, "rfq_nmbranch")
		rfq_nmvendor, _ := jsonparser.GetString(value, "rfq_nmvendor")
		rfq_totalitem, _ := jsonparser.GetFloat(value, "rfq_totalitem")
		rfq_totalrfq, _ := jsonparser.GetFloat(value, "rfq_totalrfq")
		rfq_status, _ := jsonparser.GetString(value, "rfq_status")
		rfq_status_css, _ := jsonparser.GetString(value, "rfq_status_css")
		rfq_create, _ := jsonparser.GetString(value, "rfq_create")
		rfq_update, _ := jsonparser.GetString(value, "rfq_update")

		obj.Rfq_id = rfq_id
		obj.Rfq_date = rfq_date
		obj.Rfq_idbranch = rfq_idbranch
		obj.Rfq_idvendor = rfq_idvendor
		obj.Rfq_idcurr = rfq_idcurr
		obj.Rfq_tipedoc = rfq_tipedoc
		obj.Rfq_nmbranch = rfq_nmbranch
		obj.Rfq_nmvendor = rfq_nmvendor
		obj.Rfq_totalitem = float64(rfq_totalitem)
		obj.Rfq_totalrfq = float64(rfq_totalrfq)
		obj.Rfq_status = rfq_status
		obj.Rfq_status_css = rfq_status_css
		obj.Rfq_create = rfq_create
		obj.Rfq_update = rfq_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(Listbranch_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		branch_id, _ := jsonparser.GetString(value, "branch_id")
		branch_name, _ := jsonparser.GetString(value, "branch_name")

		objbranch.Branch_id = branch_id
		objbranch.Branch_name = branch_name
		arraobjbranch = append(arraobjbranch, objbranch)
	})
	jsonparser.ArrayEach(Listcurr_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		curr_id, _ := jsonparser.GetString(value, "curr_id")

		objcurr.Curr_id = curr_id
		arraobjcurr = append(arraobjcurr, objcurr)
	})
	if !flag {
		result, err := models.Fetch_rfqHome(client.Rfq_search, client.Rfq_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldrfq_home_redis+"_"+strconv.Itoa(client.Rfq_page)+"_"+client.Rfq_search, result, 60*time.Minute)
		fmt.Println("RFQ MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("RFQ CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"listbranch":  arraobjbranch,
			"listcurr":    arraobjcurr,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"time":        time.Since(render_page).String(),
		})
	}
}
func Rfqdetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_rfqdetail)
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

	var obj entities.Model_rfqdetail
	var arraobj []entities.Model_rfqdetail
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldrfq_home_redis + "_" + client.Rfq_id)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		rfqdetail_id, _ := jsonparser.GetString(value, "rfqdetail_id")
		rfqdetail_idpurchaserequestdetail, _ := jsonparser.GetString(value, "purchaserequestdetail_idpurchaserequest")
		rfqdetail_idpurchaserequest, _ := jsonparser.GetString(value, "rfqdetail_idpurchaserequest")
		rfqdetail_nmdepartement, _ := jsonparser.GetString(value, "rfqdetail_nmdepartement")
		rfqdetail_nmemployee, _ := jsonparser.GetString(value, "rfqdetail_nmemployee")
		rfqdetail_iditem, _ := jsonparser.GetString(value, "rfqdetail_iditem")
		rfqdetail_nmitem, _ := jsonparser.GetString(value, "rfqdetail_nmitem")
		rfqdetail_descitem, _ := jsonparser.GetString(value, "rfqdetail_descitem")
		rfqdetail_qty, _ := jsonparser.GetFloat(value, "rfqdetail_qty")
		rfqdetail_iduom, _ := jsonparser.GetString(value, "rfqdetail_iduom")
		rfqdetail_price, _ := jsonparser.GetFloat(value, "rfqdetail_price")
		rfqdetail_status, _ := jsonparser.GetString(value, "rfqdetail_status")
		rfqdetail_status_css, _ := jsonparser.GetString(value, "rfqdetail_status_css")
		rfqdetail_create, _ := jsonparser.GetString(value, "rfqdetail_create")
		rfqdetail_update, _ := jsonparser.GetString(value, "rfqdetail_update")

		obj.Rfqdetail_id = rfqdetail_id
		obj.Rfqdetail_idpurchaserequestdetail = rfqdetail_idpurchaserequestdetail
		obj.Rfqdetail_idpurchaserequest = rfqdetail_idpurchaserequest
		obj.Rfqdetail_nmdepartement = rfqdetail_nmdepartement
		obj.Rfqdetail_nmemployee = rfqdetail_nmemployee
		obj.Rfqdetail_iditem = rfqdetail_iditem
		obj.Rfqdetail_nmitem = rfqdetail_nmitem
		obj.Rfqdetail_descitem = rfqdetail_descitem
		obj.Rfqdetail_qty = float64(rfqdetail_qty)
		obj.Rfqdetail_iduom = rfqdetail_iduom
		obj.Rfqdetail_price = float64(rfqdetail_price)
		obj.Rfqdetail_status = rfqdetail_status
		obj.Rfqdetail_status_css = rfqdetail_status_css
		obj.Rfqdetail_create = rfqdetail_create
		obj.Rfqdetail_update = rfqdetail_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_rfqDetail(client.Rfq_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldrfq_home_redis+"_"+client.Rfq_id, result, 60*time.Minute)
		fmt.Println("RFQ DETAIL MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("RFQ DETAIL CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func RfqSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_rfqsave)
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

	// admin, idrecord, idbranch, idvendor, idcurr, tipedoc, listdetail, sData string, total_item, subtotalpr float32
	result, err := models.Save_rfq(
		client_admin,
		client.Rfq_id, client.Rfq_idbranch, client.Rfq_idvendor, client.Rfq_idcurr, client.Rfq_tipedoc,
		client.Rfq_listdetail, client.Sdata,
		client.Rfq_totalitem, client.Rfq_subtotal)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_rfq(client.Rfq_search, client.Rfq_id, client.Rfq_page)
	return c.JSON(result)
}
func RfqstatusSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_rfqstatus)
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

	// aadmin, idrecord, status string
	result, err := models.Save_rfqStatus(
		client_admin,
		client.Rfq_id, client.Rfq_status)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_rfq("", client.Rfq_id, 0)
	return c.JSON(result)
}
func _deleteredis_rfq(search, idrfq string, page int) {
	val_master := helpers.DeleteRedis(Fieldrfq_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND RFQ : %d\n", val_master)

	val_master_detail := helpers.DeleteRedis(Fieldrfq_home_redis + "_" + idrfq)
	fmt.Printf("Redis Delete BACKEND RFQ DETAIL : %d\n", val_master_detail)

}
