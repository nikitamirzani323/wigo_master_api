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

const Fieldpurchaserequest_home_redis = "PURCAHSEREQUEST_BACKEND"

func Purchaserequesthome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_purchaserequest)
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

	if client.Purchaserequest_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldpurchaserequest_home_redis + "_" + strconv.Itoa(client.Purchaserequest_page) + "_" + client.Purchaserequest_search)
		fmt.Printf("Redis Delete BACKEND DEPARTEMENT : %d", val_pattern)
	}

	var obj entities.Model_purchaserequest
	var arraobj []entities.Model_purchaserequest
	var objbranch entities.Model_branchshare
	var arraobjbranch []entities.Model_branchshare
	var objdepartement entities.Model_departementshare
	var arraobjdepartement []entities.Model_departementshare
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldpurchaserequest_home_redis + "_" + strconv.Itoa(client.Purchaserequest_page) + "_" + client.Purchaserequest_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	Listbranch_RD, _, _, _ := jsonparser.Get(jsonredis, "listbranch")
	Listdepartement_RD, _, _, _ := jsonparser.Get(jsonredis, "listdepartement")
	Listcurr_RD, _, _, _ := jsonparser.Get(jsonredis, "listcurr")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		purchaserequest_id, _ := jsonparser.GetString(value, "purchaserequest_id")
		purchaserequest_date, _ := jsonparser.GetString(value, "purchaserequest_date")
		purchaserequest_idbranch, _ := jsonparser.GetString(value, "purchaserequest_idbranch")
		purchaserequest_iddepartement, _ := jsonparser.GetString(value, "purchaserequest_iddepartement")
		purchaserequest_idemployee, _ := jsonparser.GetString(value, "purchaserequest_idemployee")
		purchaserequest_idcurr, _ := jsonparser.GetString(value, "purchaserequest_idcurr")
		purchaserequest_tipedoc, _ := jsonparser.GetString(value, "purchaserequest_tipedoc")
		purchaserequest_periodedoc, _ := jsonparser.GetString(value, "purchaserequest_periodedoc")
		purchaserequest_nmbranch, _ := jsonparser.GetString(value, "purchaserequest_nmbranch")
		purchaserequest_nmdepartement, _ := jsonparser.GetString(value, "purchaserequest_nmdepartement")
		purchaserequest_nmemployee, _ := jsonparser.GetString(value, "purchaserequest_nmemployee")
		purchaserequest_totalitem, _ := jsonparser.GetFloat(value, "purchaserequest_totalitem")
		purchaserequest_totalpr, _ := jsonparser.GetFloat(value, "purchaserequest_totalpr")
		purchaserequest_totalpo, _ := jsonparser.GetFloat(value, "purchaserequest_totalpo")
		purchaserequest_remark, _ := jsonparser.GetString(value, "purchaserequest_remark")
		purchaserequest_docexpire, _ := jsonparser.GetString(value, "purchaserequest_docexpire")
		purchaserequest_status, _ := jsonparser.GetString(value, "purchaserequest_status")
		purchaserequest_status_css, _ := jsonparser.GetString(value, "purchaserequest_status_css")
		purchaserequest_create, _ := jsonparser.GetString(value, "purchaserequest_create")
		purchaserequest_update, _ := jsonparser.GetString(value, "purchaserequest_update")

		obj.Purchaserequest_id = purchaserequest_id
		obj.Purchaserequest_date = purchaserequest_date
		obj.Purchaserequest_idbranch = purchaserequest_idbranch
		obj.Purchaserequest_iddepartement = purchaserequest_iddepartement
		obj.Purchaserequest_idemployee = purchaserequest_idemployee
		obj.Purchaserequest_idcurr = purchaserequest_idcurr
		obj.Purchaserequest_tipedoc = purchaserequest_tipedoc
		obj.Purchaserequest_periodedoc = purchaserequest_periodedoc
		obj.Purchaserequest_nmbranch = purchaserequest_nmbranch
		obj.Purchaserequest_nmdepartement = purchaserequest_nmdepartement
		obj.Purchaserequest_nmemployee = purchaserequest_nmemployee
		obj.Purchaserequest_totalitem = float64(purchaserequest_totalitem)
		obj.Purchaserequest_totalpr = float64(purchaserequest_totalpr)
		obj.Purchaserequest_totalpo = float64(purchaserequest_totalpo)
		obj.Purchaserequest_remark = purchaserequest_remark
		obj.Purchaserequest_docexpire = purchaserequest_docexpire
		obj.Purchaserequest_status = purchaserequest_status
		obj.Purchaserequest_status_css = purchaserequest_status_css
		obj.Purchaserequest_create = purchaserequest_create
		obj.Purchaserequest_update = purchaserequest_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(Listbranch_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		branch_id, _ := jsonparser.GetString(value, "branch_id")
		branch_name, _ := jsonparser.GetString(value, "branch_name")

		objbranch.Branch_id = branch_id
		objbranch.Branch_name = branch_name
		arraobjbranch = append(arraobjbranch, objbranch)
	})
	jsonparser.ArrayEach(Listdepartement_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		departement_id, _ := jsonparser.GetString(value, "departement_id")
		departement_name, _ := jsonparser.GetString(value, "departement_name")

		objdepartement.Departement_id = departement_id
		objdepartement.Departement_name = departement_name
		arraobjdepartement = append(arraobjdepartement, objdepartement)
	})
	jsonparser.ArrayEach(Listcurr_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		curr_id, _ := jsonparser.GetString(value, "curr_id")

		objcurr.Curr_id = curr_id
		arraobjcurr = append(arraobjcurr, objcurr)
	})
	if !flag {
		result, err := models.Fetch_purchaserequestHome(client.Purchaserequest_search, client.Purchaserequest_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldpurchaserequest_home_redis+"_"+strconv.Itoa(client.Purchaserequest_page)+"_"+client.Purchaserequest_search, result, 60*time.Minute)
		fmt.Println("PURCHASE REQUEST MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("PURCHASE REQUEST CACHE")
		return c.JSON(fiber.Map{
			"status":          fiber.StatusOK,
			"message":         "Success",
			"record":          arraobj,
			"listbranch":      arraobjbranch,
			"listdepartement": arraobjdepartement,
			"listcurr":        arraobjcurr,
			"perpage":         perpage_RD,
			"totalrecord":     totalrecord_RD,
			"time":            time.Since(render_page).String(),
		})
	}
}
func Purchaserequestdetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_purchaserequestdetail)
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

	var obj entities.Model_purchaserequestdetail
	var arraobj []entities.Model_purchaserequestdetail
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldpurchaserequest_home_redis + "_" + client.Purchaserequest_id)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		purchaserequestdetail_id, _ := jsonparser.GetString(value, "purchaserequestdetail_id")
		purchaserequestdetail_idpurchaserequest, _ := jsonparser.GetString(value, "purchaserequestdetail_idpurchaserequest")
		purchaserequestdetail_iditem, _ := jsonparser.GetString(value, "purchaserequestdetail_iditem")
		purchaserequestdetail_nmitem, _ := jsonparser.GetString(value, "purchaserequestdetail_nmitem")
		purchaserequestdetail_descitem, _ := jsonparser.GetString(value, "purchaserequestdetail_descitem")
		purchaserequestdetail_purpose, _ := jsonparser.GetString(value, "purchaserequestdetail_purpose")
		purchaserequestdetail_qty, _ := jsonparser.GetFloat(value, "purchaserequestdetail_qty")
		purchaserequestdetail_iduom, _ := jsonparser.GetString(value, "purchaserequestdetail_iduom")
		purchaserequestdetail_price, _ := jsonparser.GetFloat(value, "purchaserequestdetail_price")
		purchaserequestdetail_status, _ := jsonparser.GetString(value, "purchaserequestdetail_status")
		purchaserequestdetail_status_css, _ := jsonparser.GetString(value, "purchaserequestdetail_status_css")
		purchaserequestdetail_create, _ := jsonparser.GetString(value, "purchaserequestdetail_create")
		purchaserequestdetail_update, _ := jsonparser.GetString(value, "purchaserequestdetail_update")

		obj.Purchaserequestdetail_id = purchaserequestdetail_id
		obj.Purchaserequestdetail_idpurchaserequest = purchaserequestdetail_idpurchaserequest
		obj.Purchaserequestdetail_iditem = purchaserequestdetail_iditem
		obj.Purchaserequestdetail_nmitem = purchaserequestdetail_nmitem
		obj.Purchaserequestdetail_descitem = purchaserequestdetail_descitem
		obj.Purchaserequestdetail_purpose = purchaserequestdetail_purpose
		obj.Purchaserequestdetail_qty = float32(purchaserequestdetail_qty)
		obj.Purchaserequestdetail_iduom = purchaserequestdetail_iduom
		obj.Purchaserequestdetail_price = float32(purchaserequestdetail_price)
		obj.Purchaserequestdetail_status = purchaserequestdetail_status
		obj.Purchaserequestdetail_status_css = purchaserequestdetail_status_css
		obj.Purchaserequestdetail_create = purchaserequestdetail_create
		obj.Purchaserequestdetail_update = purchaserequestdetail_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_purchaserequestDetail(client.Purchaserequest_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldpurchaserequest_home_redis+"_"+client.Purchaserequest_id, result, 60*time.Minute)
		fmt.Println("PURCHASE REQUEST DETAIL MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("PURCHASE REQUEST DETAIL CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Purchaserequestdetailview(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_prdetail_view)
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

	var obj entities.Model_prdetail_view
	var arraobj []entities.Model_prdetail_view
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldpurchaserequest_home_redis + "_" + client.Purchaserequest_tipedoc + "_" + client.Purchaserequest_idbranch)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		prdetailview_id, _ := jsonparser.GetString(value, "prdetailview_id")
		prdetailview_idpurchaserequest, _ := jsonparser.GetString(value, "prdetailview_idpurchaserequest")
		prdetailview_date, _ := jsonparser.GetString(value, "prdetailview_date")
		prdetailview_tipedoc, _ := jsonparser.GetString(value, "prdetailview_tipedoc")
		prdetailview_nmbranch, _ := jsonparser.GetString(value, "prdetailview_nmbranch")
		prdetailview_nmdepartement, _ := jsonparser.GetString(value, "prdetailview_nmdepartement")
		prdetailview_nmemployee, _ := jsonparser.GetString(value, "prdetailview_nmemployee")
		prdetailview_iditem, _ := jsonparser.GetString(value, "prdetailview_iditem")
		prdetailview_nmitem, _ := jsonparser.GetString(value, "prdetailview_nmitem")
		prdetailview_descitem, _ := jsonparser.GetString(value, "prdetailview_descitem")
		prdetailview_purpose, _ := jsonparser.GetString(value, "prdetailview_purpose")
		prdetailview_qty, _ := jsonparser.GetFloat(value, "prdetailview_qty")
		prdetailview_qty_po, _ := jsonparser.GetFloat(value, "prdetailview_qty_po")
		prdetailview_iduom, _ := jsonparser.GetString(value, "prdetailview_iduom")
		prdetailview_price, _ := jsonparser.GetFloat(value, "prdetailview_price")
		prdetailview_status, _ := jsonparser.GetString(value, "prdetailview_status")
		prdetailview_status_css, _ := jsonparser.GetString(value, "prdetailview_status_css")

		obj.Prdetailview_id = prdetailview_id
		obj.Prdetailview_idpurchaserequest = prdetailview_idpurchaserequest
		obj.Prdetailview_date = prdetailview_date
		obj.Prdetailview_tipedoc = prdetailview_tipedoc
		obj.Prdetailview_nmbranch = prdetailview_nmbranch
		obj.Prdetailview_nmdepartement = prdetailview_nmdepartement
		obj.Prdetailview_nmemployee = prdetailview_nmemployee
		obj.Prdetailview_iditem = prdetailview_iditem
		obj.Prdetailview_nmitem = prdetailview_nmitem
		obj.Prdetailview_descitem = prdetailview_descitem
		obj.Prdetailview_purpose = prdetailview_purpose
		obj.Prdetailview_qty = float32(prdetailview_qty)
		obj.Prdetailview_qty_po = float32(prdetailview_qty_po)
		obj.Prdetailview_iduom = prdetailview_iduom
		obj.Prdetailview_price = float32(prdetailview_price)
		obj.Prdetailview_status = prdetailview_status
		obj.Prdetailview_status_css = prdetailview_status_css

		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_prdetail_view(client.Purchaserequest_idbranch, client.Purchaserequest_tipedoc)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldpurchaserequest_home_redis+"_"+client.Purchaserequest_tipedoc+"_"+client.Purchaserequest_idbranch, result, 60*time.Minute)
		fmt.Println("PURCHASE REQUEST DETAIL VIEW MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("PURCHASE REQUEST DETAIL VIEW CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PurchaserequestSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_purchaserequestsave)
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

	// admin, idrecord, idbranch, iddepartement, idemployee, idcurr, tipedoc, remark, listdetail, sData string, total_item, subtotalpr float32
	result, err := models.Save_purchaserequest(
		client_admin,
		client.Purchaserequest_id, client.Purchaserequest_idbranch, client.Purchaserequest_iddepartement, client.Purchaserequest_idemployee, client.Purchaserequest_idcurr,
		client.Purchaserequest_tipedoc, client.Purchaserequest_remark, client.Purchaserequest_listdetail, client.Sdata,
		client.Purchaserequest_totalitem, client.Purchaserequest_subtotal)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_purchaserequest(client.Purchaserequest_search, "", client.Purchaserequest_page)
	return c.JSON(result)
}
func PurchaserequeststatusSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_purchaserequeststatus)
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
	result, err := models.Save_purchaserequestStatus(
		client_admin,
		client.Purchaserequest_id, client.Purchaserequest_status)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_purchaserequest("", client.Purchaserequest_id, 0)
	return c.JSON(result)
}
func _deleteredis_purchaserequest(search, idpurchase string, page int) {
	val_master := helpers.DeleteRedis(Fieldpurchaserequest_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND PURCHASE REQUEST : %d\n", val_master)

	val_master_detail := helpers.DeleteRedis(Fieldpurchaserequest_home_redis + "_" + idpurchase)
	fmt.Printf("Redis Delete BACKEND PURCHASE REQUEST DETAIL : %d\n", val_master_detail)

}
