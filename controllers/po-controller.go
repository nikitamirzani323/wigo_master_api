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

const FieldPo_home_redis = "PO_BACKEND"

func Pohome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_po)
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

	if client.Po_search != "" {
		val_pattern := helpers.DeleteRedis(FieldPo_home_redis + "_" + strconv.Itoa(client.Po_page) + "_" + client.Po_search)
		fmt.Printf("Redis Delete BACKEND RFQ : %d", val_pattern)
	}

	var obj entities.Model_po
	var arraobj []entities.Model_po
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldrfq_home_redis + "_" + strconv.Itoa(client.Po_page) + "_" + client.Po_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		po_id, _ := jsonparser.GetString(value, "po_id")
		po_date, _ := jsonparser.GetString(value, "po_date")
		po_idrfq, _ := jsonparser.GetString(value, "po_idrfq")
		po_idbranch, _ := jsonparser.GetString(value, "po_idbranch")
		po_idvendor, _ := jsonparser.GetString(value, "po_idvendor")
		po_idcurr, _ := jsonparser.GetString(value, "po_idcurr")
		po_tipedoc, _ := jsonparser.GetString(value, "po_tipedoc")
		po_nmbranch, _ := jsonparser.GetString(value, "po_nmbranch")
		po_nmvendor, _ := jsonparser.GetString(value, "po_nmvendor")
		po_discount, _ := jsonparser.GetFloat(value, "po_discount")
		po_ppn, _ := jsonparser.GetFloat(value, "po_ppn")
		po_pph, _ := jsonparser.GetFloat(value, "po_pph")
		po_totalitem, _ := jsonparser.GetFloat(value, "po_totalitem")
		po_subtotal, _ := jsonparser.GetFloat(value, "po_subtotal")
		po_grandtotal, _ := jsonparser.GetFloat(value, "po_grandtotal")
		po_status, _ := jsonparser.GetString(value, "po_status")
		po_status_css, _ := jsonparser.GetString(value, "po_status_css")
		po_create, _ := jsonparser.GetString(value, "po_create")
		po_update, _ := jsonparser.GetString(value, "po_update")

		obj.Po_id = po_id
		obj.Po_date = po_date
		obj.Po_idrfq = po_idrfq
		obj.Po_idbranch = po_idbranch
		obj.Po_idvendor = po_idvendor
		obj.Po_idcurr = po_idcurr
		obj.Po_tipedoc = po_tipedoc
		obj.Po_nmbranch = po_nmbranch
		obj.Po_nmvendor = po_nmvendor
		obj.Po_discount = float64(po_discount)
		obj.Po_ppn = float64(po_ppn)
		obj.Po_pph = float64(po_pph)
		obj.Po_totalitem = float64(po_totalitem)
		obj.Po_subtotal = float64(po_subtotal)
		obj.Po_grandtotal = float64(po_grandtotal)
		obj.Po_status = po_status
		obj.Po_status_css = po_status_css
		obj.Po_create = po_create
		obj.Po_update = po_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_poHome(client.Po_search, client.Po_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(FieldPo_home_redis+"_"+strconv.Itoa(client.Po_page)+"_"+client.Po_search, result, 60*time.Minute)
		fmt.Println("PO MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("PO CACHE")
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
func Podetail(c *fiber.Ctx) error {
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
func PoSave(c *fiber.Ctx) error {
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

	_deleteredis_po(client.Rfq_search, client.Rfq_id, client.Rfq_page)
	return c.JSON(result)
}
func PostatusSave(c *fiber.Ctx) error {
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

	_deleteredis_po("", client.Rfq_id, 0)
	return c.JSON(result)
}
func _deleteredis_po(search, idrfq string, page int) {
	val_master := helpers.DeleteRedis(Fieldrfq_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND RFQ : %d\n", val_master)

	val_master_detail := helpers.DeleteRedis(Fieldrfq_home_redis + "_" + idrfq)
	fmt.Printf("Redis Delete BACKEND RFQ DETAIL : %d\n", val_master_detail)

}
