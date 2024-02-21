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

const Fieldmerek_home_redis = "LISTMEREK_BACKEND"
const Fieldcateitem_home_redis = "LISTCATEITEM_BACKEND"
const Fielditem_home_redis = "LISTITEM_BACKEND"
const Fielditem_share_redis = "LISTITEM_SHARE_BACKEND"
const Fielditemuom_home_redis = "LISTITEMUOM_BACKEND"

func Merekhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_merek)
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
	if client.Merek_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldmerek_home_redis + "_" + strconv.Itoa(client.Merek_page) + "_" + client.Merek_search)
		fmt.Printf("Redis Delete BACKEND MEREK : %d", val_pattern)
	}

	var obj entities.Model_merek
	var arraobj []entities.Model_merek
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmerek_home_redis + "_" + strconv.Itoa(client.Merek_page) + "_" + client.Merek_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		merek_id, _ := jsonparser.GetInt(value, "merek_id")
		merek_name, _ := jsonparser.GetString(value, "merek_name")
		merek_status, _ := jsonparser.GetString(value, "merek_status")
		merek_status_css, _ := jsonparser.GetString(value, "merek_status_css")
		merek_create, _ := jsonparser.GetString(value, "merek_create")
		merek_update, _ := jsonparser.GetString(value, "merek_update")

		obj.Merek_id = int(merek_id)
		obj.Merek_name = merek_name
		obj.Merek_status = merek_status
		obj.Merek_status_css = merek_status_css
		obj.Merek_create = merek_create
		obj.Merek_update = merek_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_merekHome(client.Merek_search, client.Merek_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmerek_home_redis+"_"+strconv.Itoa(client.Merek_page)+"_"+client.Merek_search, result, 60*time.Minute)
		fmt.Println("MEREK MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("MEREK CACHE")
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
func Merekshare(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_merek)
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
	if client.Merek_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldmerek_home_redis + "_" + client.Merek_search)
		fmt.Printf("Redis Delete BACKEND MEREK : %d", val_pattern)
	}

	var obj entities.Model_merekshare
	var arraobj []entities.Model_merekshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmerek_home_redis + "_" + client.Merek_search)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		merek_id, _ := jsonparser.GetInt(value, "merek_id")
		merek_name, _ := jsonparser.GetString(value, "merek_name")

		obj.Merek_id = int(merek_id)
		obj.Merek_name = merek_name
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_merekHome(client.Merek_search, client.Merek_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmerek_home_redis+"_"+client.Merek_search, result, 60*time.Minute)
		fmt.Println("MEREK SHARE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("MEREK SHARE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Cateitemhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_cateitem)
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
	fmt.Println(client.Cateitem_page)
	if client.Cateitem_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldcateitem_home_redis + "_" + strconv.Itoa(client.Cateitem_page) + "_" + client.Cateitem_search)
		fmt.Printf("Redis Delete BACKEND CATEITEM : %d", val_pattern)
	}

	var obj entities.Model_cateitem
	var arraobj []entities.Model_cateitem
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcateitem_home_redis + "_" + strconv.Itoa(client.Cateitem_page) + "_" + client.Cateitem_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		cateitem_id, _ := jsonparser.GetInt(value, "cateitem_id")
		cateitem_name, _ := jsonparser.GetString(value, "cateitem_name")
		cateitem_status, _ := jsonparser.GetString(value, "cateitem_status")
		cateitem_status_css, _ := jsonparser.GetString(value, "cateitem_status_css")
		cateitem_create, _ := jsonparser.GetString(value, "cateitem_create")
		cateitem_update, _ := jsonparser.GetString(value, "cateitem_update")

		obj.Cateitem_id = int(cateitem_id)
		obj.Cateitem_name = cateitem_name
		obj.Cateitem_status = cateitem_status
		obj.Cateitem_status_css = cateitem_status_css
		obj.Cateitem_create = cateitem_create
		obj.Cateitem_update = cateitem_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_catetemHome(client.Cateitem_search, client.Cateitem_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcateitem_home_redis+"_"+strconv.Itoa(client.Cateitem_page)+"_"+client.Cateitem_search, result, 60*time.Minute)
		fmt.Println("CATE ITEM MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("CATE ITEM CACHE")
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
func Itemhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_item)
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
	fmt.Println(client.Item_page)
	if client.Item_search != "" {
		val_pattern := helpers.DeleteRedis(Fielditem_home_redis + "_" + strconv.Itoa(client.Item_page) + "_" + client.Item_search)
		fmt.Printf("Redis Delete BACKEND CATEITEM : %d", val_pattern)
	}

	var obj entities.Model_item
	var arraobj []entities.Model_item
	var objcateitem entities.Model_cateitemshare
	var arraobjcateitem []entities.Model_cateitemshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielditem_home_redis + "_" + strconv.Itoa(client.Item_page) + "_" + client.Item_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listcateitem_RD, _, _, _ := jsonparser.Get(jsonredis, "listcateitem")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		item_id, _ := jsonparser.GetString(value, "item_id")
		item_idmerek, _ := jsonparser.GetInt(value, "item_idmerek")
		item_nmmerek, _ := jsonparser.GetString(value, "item_nmmerek")
		item_idcateitem, _ := jsonparser.GetInt(value, "item_idcateitem")
		item_nmcateitem, _ := jsonparser.GetString(value, "item_nmcateitem")
		item_iduom, _ := jsonparser.GetString(value, "item_iduom")
		item_name, _ := jsonparser.GetString(value, "item_name")
		item_descp, _ := jsonparser.GetString(value, "item_descp")
		item_urlimg, _ := jsonparser.GetString(value, "item_urlimg")
		item_inventory, _ := jsonparser.GetString(value, "item_inventory")
		item_sales, _ := jsonparser.GetString(value, "item_sales")
		item_purchase, _ := jsonparser.GetString(value, "item_purchase")
		item_inventory_css, _ := jsonparser.GetString(value, "item_inventory_css")
		item_sales_css, _ := jsonparser.GetString(value, "item_sales_css")
		item_purchase_css, _ := jsonparser.GetString(value, "item_purchase_css")
		item_status, _ := jsonparser.GetString(value, "item_status")
		item_status_css, _ := jsonparser.GetString(value, "item_status_css")
		item_create, _ := jsonparser.GetString(value, "item_create")
		item_update, _ := jsonparser.GetString(value, "item_update")

		obj.Item_id = item_id
		obj.Item_idmerek = int(item_idmerek)
		obj.Item_nmmerek = item_nmmerek
		obj.Item_idcateitem = int(item_idcateitem)
		obj.Item_nmcateitem = item_nmcateitem
		obj.Item_iduom = item_iduom
		obj.Item_name = item_name
		obj.Item_descp = item_descp
		obj.Item_urlimg = item_urlimg
		obj.Item_inventory = item_inventory
		obj.Item_sales = item_sales
		obj.Item_purchase = item_purchase
		obj.Item_inventory_css = item_inventory_css
		obj.Item_purchase_css = item_purchase_css
		obj.Item_sales_css = item_sales_css
		obj.Item_status = item_status
		obj.Item_status_css = item_status_css
		obj.Item_create = item_create
		obj.Item_update = item_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcateitem_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		cateitem_id, _ := jsonparser.GetInt(value, "cateitem_id")
		cateitem_name, _ := jsonparser.GetString(value, "cateitem_name")

		objcateitem.Cateitem_id = int(cateitem_id)
		objcateitem.Cateitem_name = cateitem_name
		arraobjcateitem = append(arraobjcateitem, objcateitem)
	})
	if !flag {
		result, err := models.Fetch_itemHome(client.Item_search, client.Item_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fielditem_home_redis+"_"+strconv.Itoa(client.Item_page)+"_"+client.Item_search, result, 60*time.Minute)
		fmt.Println("ITEM MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("ITEM CACHE")
		return c.JSON(fiber.Map{
			"status":       fiber.StatusOK,
			"message":      "Success",
			"record":       arraobj,
			"listcateitem": arraobjcateitem,
			"perpage":      perpage_RD,
			"totalrecord":  totalrecord_RD,
			"time":         time.Since(render_page).String(),
		})
	}
}
func Itemshare(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_item)
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

	var obj entities.Model_itemshare
	var arraobj []entities.Model_itemshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielditem_share_redis + "_" + client.Item_search)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		itemshare_id, _ := jsonparser.GetString(value, "itemshare_id")
		itemshare_nmcateitem, _ := jsonparser.GetString(value, "itemshare_nmcateitem")
		itemshare_name, _ := jsonparser.GetString(value, "itemshare_name")
		itemshare_descp, _ := jsonparser.GetString(value, "itemshare_descp")
		itemshare_urlimg, _ := jsonparser.GetString(value, "itemshare_urlimg")

		var objitemuom entities.Model_itemuomshare
		var arraobjitemuom []entities.Model_itemuomshare
		record_itemuom_RD, _, _, _ := jsonparser.Get(value, "itemshare_uom")
		jsonparser.ArrayEach(record_itemuom_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			itemuom_iduom, _ := jsonparser.GetString(value, "itemuom_iduom")

			objitemuom.Itemuom_iduom = itemuom_iduom
			arraobjitemuom = append(arraobjitemuom, objitemuom)
		})

		obj.Itemshare_id = itemshare_id
		obj.Itemshare_nmcateitem = itemshare_nmcateitem
		obj.Itemshare_name = itemshare_name
		obj.Itemshare_descp = itemshare_descp
		obj.Itemshare_urlimg = itemshare_urlimg
		obj.Itemshare_uom = arraobjitemuom

		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_itemShare(client.Item_search)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fielditem_share_redis+"_"+client.Item_search, result, 30*time.Minute)
		fmt.Println("ITEM MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("ITEM CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Itemuom(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_itemuom)
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

	var obj entities.Model_itemuom
	var arraobj []entities.Model_itemuom
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielditemuom_home_redis + "_" + client.Item_id)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		itemuom_id, _ := jsonparser.GetInt(value, "itemuom_id")
		itemuom_iduom, _ := jsonparser.GetString(value, "itemuom_iduom")
		itemuom_nmuom, _ := jsonparser.GetString(value, "itemuom_nmuom")
		itemuom_default, _ := jsonparser.GetString(value, "itemuom_default")
		itemuom_default_css, _ := jsonparser.GetString(value, "itemuom_default_css")
		itemuom_conversion, _ := jsonparser.GetFloat(value, "itemuom_conversion")
		itemuom_create, _ := jsonparser.GetString(value, "itemuom_create")
		itemuom_update, _ := jsonparser.GetString(value, "itemuom_update")

		obj.Itemuom_id = int(itemuom_id)
		obj.Itemuom_iduom = itemuom_iduom
		obj.Itemuom_nmuom = itemuom_nmuom
		obj.Itemuom_default = itemuom_default
		obj.Itemuom_default_css = itemuom_default_css
		obj.Itemuom_conversion = float32(itemuom_conversion)
		obj.Itemuom_create = itemuom_create
		obj.Itemuom_update = itemuom_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_itemuom(client.Item_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fielditemuom_home_redis+"_"+client.Item_id, result, 60*time.Minute)
		fmt.Println("ITEM UOM MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("ITEM UOM CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func MerekSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_mereksave)
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
	result, err := models.Save_merek(
		client_admin,
		client.Merek_name, client.Merek_status, client.Sdata, client.Merek_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_item(client.Merek_search, "", client.Merek_page)
	return c.JSON(result)
}
func CateitemSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_cateitemsave)
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
	result, err := models.Save_cateitem(
		client_admin,
		client.Cateitem_name, client.Cateitem_status, client.Sdata, client.Cateitem_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_item(client.Cateitem_search, "", client.Cateitem_page)
	return c.JSON(result)
}
func ItemSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_itemsave)
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

	// admin, idrecord, iduom, name, descp, urlimgitem, inventory, sales, purchase, status, sData string, idcateitem int
	result, err := models.Save_item(
		client_admin,
		client.Item_id, client.Item_iduom, client.Item_name, client.Item_descp, client.Item_urlimg,
		client.Item_inventory, client.Item_sales, client.Item_purchase, client.Item_status,
		client.Sdata, client.Item_idcateitem, client.Item_idmerek)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_item(client.Item_search, client.Item_id, client.Item_page)
	return c.JSON(result)
}
func ItemuomSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_itemuomsave)
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

	// admin, iditem, iduom, default_iduom, sData string, idrecord int, convertion float32
	result, err := models.Save_itemuom(
		client_admin,
		client.Itemuom_iditem, client.Itemuom_iduom, client.Itemuom_default,
		client.Sdata, client.Itemuom_id, client.Itemuom_conversion)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_item(client.Itemuom_search, client.Itemuom_iditem, client.Itemuom_page)
	return c.JSON(result)
}
func ItemuomDelete(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_itemuomdelete)
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
	// user := c.Locals("jwt").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)
	// temp_decp := helpers.Decryption(name)
	// _, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Delete_itemuom(client.Itemuom_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_item(client.Itemuom_search, client.Itemuom_iditem, client.Itemuom_page)
	return c.JSON(result)
}
func _deleteredis_item(search, iditem string, page int) {
	val_master_merek := helpers.DeleteRedis(Fieldmerek_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND MEREK : %d\n", val_master_merek)

	val_master := helpers.DeleteRedis(Fieldcateitem_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND CATEITEM : %d\n", val_master)

	val_master_item := helpers.DeleteRedis(Fielditem_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND ITEM : %d\n", val_master_item)

	val_master_itemuom := helpers.DeleteRedis(Fielditemuom_home_redis + "_" + iditem)
	fmt.Printf("Redis Delete BACKEND ITEMUOM - %s : %d\n", iditem, val_master_itemuom)

	val_master_itemshare := helpers.DeleteRedis(Fielditem_share_redis + "_" + search)
	fmt.Printf("Redis Delete BACKEND ITEMSHARE - %s : %d\n", iditem, val_master_itemshare)
}
