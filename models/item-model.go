package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/db"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nleeper/goment"
)

const database_merek_local = configs.DB_tbl_mst_merek
const database_cateitem_local = configs.DB_tbl_mst_categoryitem
const database_item_local = configs.DB_tbl_mst_item

func Fetch_merekHome(search string, page int) (helpers.Responsepaging, error) {
	var obj entities.Model_merek
	var arraobj []entities.Model_merek
	var res helpers.Responsepaging
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := configs.PAGING_PAGE
	totalrecord := 0
	offset := page
	sql_selectcount := ""
	sql_selectcount += ""
	sql_selectcount += "SELECT "
	sql_selectcount += "COUNT(idmerek) as totalcateitem  "
	sql_selectcount += "FROM " + database_merek_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(nmmerek) LIKE '%" + strings.ToLower(search) + "%' "
	}
	row_selectcount := con.QueryRowContext(ctx, sql_selectcount)
	switch e_selectcount := row_selectcount.Scan(&totalrecord); e_selectcount {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e_selectcount)
	}

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idmerek , nmmerek, statusmerek, "
	sql_select += "createmerek, to_char(COALESCE(createdatemerek,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "updatemerek, to_char(COALESCE(updatedatemerek,now()), 'YYYY-MM-DD HH24:MI:SS') "
	sql_select += "FROM " + database_merek_local + "   "
	if search == "" {
		sql_select += "ORDER BY createdatemerek DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(nmmerek) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY createdatemerek DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idmerek_db                                                             int
			nmmerek_db, statusmerek_db                                             string
			createmerek_db, createdatemerek_db, updatemerek_db, updatedatemerek_db string
		)

		err = row.Scan(&idmerek_db, &nmmerek_db, &statusmerek_db,
			&createmerek_db, &createdatemerek_db, &updatemerek_db, &updatedatemerek_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createmerek_db != "" {
			create = createmerek_db + ", " + createdatemerek_db
		}
		if updatemerek_db != "" {
			update = updatemerek_db + ", " + updatedatemerek_db
		}
		if statusmerek_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Merek_id = idmerek_db
		obj.Merek_name = nmmerek_db
		obj.Merek_status = statusmerek_db
		obj.Merek_status_css = status_css
		obj.Merek_create = create
		obj.Merek_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_merekShare(search string) (helpers.Response, error) {
	var obj entities.Model_merekshare
	var arraobj []entities.Model_merekshare
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idmerek , nmmerek "
	sql_select += "FROM " + database_merek_local + "   "
	sql_select += "WHERE statusmerek='Y' "
	if search != "" {
		sql_select += "AND LOWER(nmmerek) LIKE '%" + strings.ToLower(search) + "%' "
	} else {
		sql_select += "ORDER BY nmmerek DESC LIMIT " + strconv.Itoa(configs.PAGING_PAGE)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idmerek_db int
			nmmerek_db string
		)

		err = row.Scan(&idmerek_db, &nmmerek_db)

		helpers.ErrorCheck(err)

		obj.Merek_id = idmerek_db
		obj.Merek_name = nmmerek_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_catetemHome(search string, page int) (helpers.Responsepaging, error) {
	var obj entities.Model_cateitem
	var arraobj []entities.Model_cateitem
	var res helpers.Responsepaging
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := configs.PAGING_PAGE
	totalrecord := 0
	offset := page
	sql_selectcount := ""
	sql_selectcount += ""
	sql_selectcount += "SELECT "
	sql_selectcount += "COUNT(idcateitem) as totalcateitem  "
	sql_selectcount += "FROM " + database_cateitem_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(nmcateitem) LIKE '%" + strings.ToLower(search) + "%' "
	}
	row_selectcount := con.QueryRowContext(ctx, sql_selectcount)
	switch e_selectcount := row_selectcount.Scan(&totalrecord); e_selectcount {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e_selectcount)
	}

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idcateitem , nmcateitem, statuscateitem, "
	sql_select += "createcateitem, to_char(COALESCE(createdatecateitem,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "updatecateitem, to_char(COALESCE(updatedatecateitem,now()), 'YYYY-MM-DD HH24:MI:SS') "
	sql_select += "FROM " + database_cateitem_local + "   "
	if search == "" {
		sql_select += "ORDER BY createdatecateitem DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(nmcateitem) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY createdatecateitem DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcateitem_db                                                                      int
			nmcateitem_db, statuscateitem_db                                                   string
			createcateitem_db, createdatecateitem_db, updatecateitem_db, updatedatecateitem_db string
		)

		err = row.Scan(&idcateitem_db, &nmcateitem_db, &statuscateitem_db,
			&createcateitem_db, &createdatecateitem_db, &updatecateitem_db, &updatedatecateitem_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createcateitem_db != "" {
			create = createcateitem_db + ", " + createdatecateitem_db
		}
		if updatecateitem_db != "" {
			update = updatecateitem_db + ", " + updatedatecateitem_db
		}
		if statuscateitem_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Cateitem_id = idcateitem_db
		obj.Cateitem_name = nmcateitem_db
		obj.Cateitem_status = statuscateitem_db
		obj.Cateitem_status_css = status_css
		obj.Cateitem_create = create
		obj.Cateitem_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_itemHome(search string, page int) (helpers.Responseitem, error) {
	var obj entities.Model_item
	var arraobj []entities.Model_item
	var objcateitem entities.Model_cateitemshare
	var arraobjcateitem []entities.Model_cateitemshare
	var res helpers.Responseitem
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := configs.PAGING_PAGE
	totalrecord := 0
	offset := page
	sql_selectcount := ""
	sql_selectcount += ""
	sql_selectcount += "SELECT "
	sql_selectcount += "COUNT(iditem) as totalitem  "
	sql_selectcount += "FROM " + database_item_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(nmitem) LIKE '%" + strings.ToLower(search) + "%' "
	}
	row_selectcount := con.QueryRowContext(ctx, sql_selectcount)
	switch e_selectcount := row_selectcount.Scan(&totalrecord); e_selectcount {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e_selectcount)
	}

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "A.iditem , A.idmerek, A.idcateitem, B.nmcateitem,  "
	sql_select += "A.nmitem , A.descpitem, A.urlimgitem, A.inventory_item, A.sales_item, A.purchase_item, A.statusitem,  "
	sql_select += "A.createitem, to_char(COALESCE(A.createdateitem,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updateitem, to_char(COALESCE(A.updatedateitem,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_item_local + "  as A  "
	sql_select += "JOIN " + database_cateitem_local + "  as B ON B.idcateitem = A.idcateitem   "
	if search == "" {
		sql_select += "WHERE LOWER(nmitem) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdateitem DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(nmcateitem) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdateitem DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idmerek_db, idcateitem_db                                                                                                           int
			iditem_db, nmcateitem_db, nmitem_db, descpitem_db, urlimgitem_db, inventory_item_db, sales_item_db, purchase_item_db, statusitem_db string
			createitem_db, createdateitem_db, updateitem_db, updatedateitem_db                                                                  string
		)

		err = row.Scan(&iditem_db, &idmerek_db, &idcateitem_db, &nmcateitem_db, &nmitem_db,
			&descpitem_db, &urlimgitem_db, &inventory_item_db, &sales_item_db, &purchase_item_db, &statusitem_db,
			&createitem_db, &createdateitem_db, &updateitem_db, &updatedateitem_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		purchase_css := configs.STATUS_CANCEL
		sales_css := configs.STATUS_CANCEL
		inventory_css := configs.STATUS_CANCEL
		if createitem_db != "" {
			create = createitem_db + ", " + createdateitem_db
		}
		if updateitem_db != "" {
			update = updateitem_db + ", " + updatedateitem_db
		}
		if statusitem_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}
		if purchase_item_db == "Y" {
			purchase_css = configs.STATUS_COMPLETE
		}
		if sales_item_db == "Y" {
			sales_css = configs.STATUS_COMPLETE
		}
		if inventory_item_db == "Y" {
			inventory_css = configs.STATUS_COMPLETE
		}

		obj.Item_id = iditem_db
		obj.Item_idmerek = idmerek_db
		obj.Item_nmmerek = _Get_merek(idmerek_db)
		obj.Item_idcateitem = idcateitem_db
		obj.Item_nmcateitem = nmcateitem_db
		obj.Item_iduom = _Get_item_uom(iditem_db)
		obj.Item_name = nmitem_db
		obj.Item_descp = descpitem_db
		obj.Item_urlimg = urlimgitem_db
		obj.Item_inventory = inventory_item_db
		obj.Item_sales = sales_item_db
		obj.Item_purchase = purchase_item_db
		obj.Item_purchase_css = purchase_css
		obj.Item_sales_css = sales_css
		obj.Item_inventory_css = inventory_css
		obj.Item_status = statusitem_db
		obj.Item_status_css = status_css
		obj.Item_create = create
		obj.Item_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	sql_selectcateitem := `SELECT 
			idcateitem , nmcateitem 
			FROM ` + database_cateitem_local + ` 
			WHERE statuscateitem = 'Y' 
			ORDER BY nmcateitem ASC    
	`
	rowcateitem, errcateitem := con.QueryContext(ctx, sql_selectcateitem)
	helpers.ErrorCheck(errcateitem)
	for rowcateitem.Next() {
		var (
			idcateitem_db int
			nmcateitem_db string
		)

		errcateitem = rowcateitem.Scan(&idcateitem_db, &nmcateitem_db)

		helpers.ErrorCheck(errcateitem)

		objcateitem.Cateitem_id = idcateitem_db
		objcateitem.Cateitem_name = nmcateitem_db
		arraobjcateitem = append(arraobjcateitem, objcateitem)
		msg = "Success"
	}
	defer rowcateitem.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listcateitem = arraobjcateitem
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_itemShare(search string) (helpers.Response, error) {
	var obj entities.Model_itemshare
	var arraobj []entities.Model_itemshare

	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "iditem , nmcateitem, nmitem,  "
	sql_select += "descpitem , urlimgitem  "
	sql_select += "FROM " + configs.DB_view_tbl_item_purchase + "  as A  "
	if search != "" {
		sql_select += "WHERE LOWER(nmcateitem) LIKE '%" + strings.ToLower(search) + "%' "
	}
	sql_select += "ORDER BY nmitem ASC   LIMIT " + strconv.Itoa(configs.PAGING_PAGE)

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			iditem_db, nmcateitem_db, nmitem_db, descpitem_db, urlimgitem_db string
		)

		err = row.Scan(&iditem_db, &nmcateitem_db, &nmitem_db, &descpitem_db, &urlimgitem_db)

		helpers.ErrorCheck(err)

		var objitemuom entities.Model_itemuomshare
		var arraobjitemuom []entities.Model_itemuomshare
		sql_itemuom := `SELECT 
				iduom  
				FROM ` + configs.DB_tbl_mst_item_uom + ` 
				WHERE iditem=$1 AND default_itemuom!='Y' 
				ORDER BY iduom ASC    
		`
		rowitemuom, erritemuom := con.QueryContext(ctx, sql_itemuom, iditem_db)
		helpers.ErrorCheck(erritemuom)
		for rowitemuom.Next() {
			var (
				iduom_db string
			)

			erritemuom = rowitemuom.Scan(&iduom_db)

			helpers.ErrorCheck(erritemuom)

			objitemuom.Itemuom_iduom = iduom_db
			arraobjitemuom = append(arraobjitemuom, objitemuom)
			msg = "Success"
		}
		defer rowitemuom.Close()

		if len(arraobjitemuom) > 0 {
			obj.Itemshare_id = iditem_db
			obj.Itemshare_nmcateitem = nmcateitem_db
			obj.Itemshare_name = nmitem_db
			obj.Itemshare_descp = descpitem_db
			obj.Itemshare_urlimg = urlimgitem_db
			obj.Itemshare_uom = arraobjitemuom
			arraobj = append(arraobj, obj)
		}

		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_itemuom(iditem string) (helpers.Response, error) {
	var obj entities.Model_itemuom
	var arraobj []entities.Model_itemuom
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "A.id_itemuom , A.iduom, B.nmuom,  "
	sql_select += "A.default_itemuom , A.conversion_itemuom,  "
	sql_select += "A.create_itemuom, to_char(COALESCE(A.createdate_itemuom,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.update_itemuom, to_char(COALESCE(A.updatedate_itemuom,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + configs.DB_tbl_mst_item_uom + "  as A  "
	sql_select += "JOIN " + configs.DB_tbl_mst_uom + "  as B ON B.iduom = A.iduom   "
	sql_select += "WHERE iditem='" + iditem + "' "
	sql_select += "ORDER BY A.default_itemuom DESC   "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			id_itemuom_db                                                                      int
			conversion_itemuom_db                                                              float32
			iduom_db, nmuom_db, default_itemuom_db                                             string
			create_itemuom_db, createdate_itemuom_db, update_itemuom_db, updatedate_itemuom_db string
		)

		err = row.Scan(&id_itemuom_db, &iduom_db, &nmuom_db, &default_itemuom_db, &conversion_itemuom_db,
			&create_itemuom_db, &createdate_itemuom_db, &update_itemuom_db, &updatedate_itemuom_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if create_itemuom_db != "" {
			create = create_itemuom_db + ", " + createdate_itemuom_db
		}
		if update_itemuom_db != "" {
			update = update_itemuom_db + ", " + updatedate_itemuom_db
		}
		if default_itemuom_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Itemuom_id = id_itemuom_db
		obj.Itemuom_iduom = iduom_db
		obj.Itemuom_nmuom = nmuom_db
		obj.Itemuom_default = default_itemuom_db
		obj.Itemuom_default_css = status_css
		obj.Itemuom_conversion = conversion_itemuom_db
		obj.Itemuom_create = create
		obj.Itemuom_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_merek(admin, name, status, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {

		sql_insert := `
				insert into
				` + database_merek_local + ` (
					idmerek , nmmerek, statusmerek,  
					createmerek, createdatemerek 
				) values (
					$1, $2, $3,   
					$4, $5 
				)
			`
		field_column := database_merek_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_merek_local, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), name, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_merek_local + `  
				SET nmmerek=$1, statusmerek=$2,  
				updatemerek=$3, updatedatemerek=$4    
				WHERE idmerek=$5   
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_merek_local, "UPDATE",
			name, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_cateitem(admin, name, status, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {

		sql_insert := `
				insert into
				` + database_cateitem_local + ` (
					idcateitem , nmcateitem, statuscateitem,  
					createcateitem, createdatecateitem 
				) values (
					$1, $2, $3,   
					$4, $5 
				)
			`
		field_column := database_cateitem_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_cateitem_local, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), name, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_cateitem_local + `  
				SET nmcateitem=$1, statuscateitem=$2,  
				updatecateitem=$3, updatedatecateitem=$4    
				WHERE idcateitem=$5   
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_cateitem_local, "UPDATE",
			name, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_item(admin, idrecord, iduom, name, descp, urlimgitem, inventory, sales, purchase, status, sData string, idcateitem, idmerek int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_item_local + ` (
					iditem , idmerek, idcateitem, nmitem,  
					descpitem , urlimgitem, inventory_item, sales_item, purchase_item, statusitem,
					createitem, createdateitem 
				) values (
					$1, $2, $3, $4,  
					$5, $6, $7, $8, $9, $10,   
					$11, $12 
				)
			`
		field_column := database_item_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := "ITEM_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_item_local, "INSERT",
			idrecord, idmerek, idcateitem, name,
			descp, urlimgitem, inventory, sales, purchase, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
			//Default UOM
			sql_insertuom := `
					insert into
					` + configs.DB_tbl_mst_item_uom + ` (
						id_itemuom , iditem, iduom,  
						default_itemuom , conversion_itemuom,
						create_itemuom, createdate_itemuom 
					) values (
						$1, $2, $3,   
						$4, $5, 
						$6, $7 
					)
				`
			field_column_uom := configs.DB_tbl_mst_item_uom + tglnow.Format("YYYY")
			idrecord_uom_counter := Get_counter(field_column_uom)
			idrecord_uom := tglnow.Format("YY") + tglnow.Format("MM") + strconv.Itoa(idrecord_uom_counter)
			flag_insertuom, msg_insertuom := Exec_SQL(sql_insertuom, configs.DB_tbl_mst_item_uom, "INSERT",
				idrecord_uom, idrecord, iduom,
				"Y", 1,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))
			if flag_insertuom {
				msg = "Succes"
			} else {
				fmt.Println(msg_insertuom)
			}

		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_item_local + `  
				SET idmerek=$1, idcateitem=$2, nmitem=$3,  
				descpitem=$4, urlimgitem=$5, inventory_item=$6, sales_item=$7, purchase_item=$8, statusitem=$9,
				updateitem=$10, updatedateitem=$11      
				WHERE iditem=$12     
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_item_local, "UPDATE",
			idmerek, idcateitem, name, descp, urlimgitem, inventory, sales, purchase, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_itemuom(admin, iditem, iduom, default_iduom, sData string, idrecord int, convertion float32) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false
	total_itemuom := 0
	if sData == "New" {
		total_itemuom = _total_item_uom(iditem)
		if total_itemuom < 4 {
			flag = CheckDBTwoField(configs.DB_tbl_mst_item_uom, "iditem", iditem, "iduom", iduom)
			if !flag {
				if default_iduom == "Y" {
					flag_reset := _reset_itemuom(admin, iditem)
					if flag_reset {
						sql_insert := `
						insert into
						` + configs.DB_tbl_mst_item_uom + ` (
							id_itemuom, iditem, iduom,  
							default_itemuom , conversion_itemuom,
							create_itemuom, createdate_itemuom 
						) values (
							$1, $2, $3,   
							$4, $5, 
							$6, $7 
						)
					`
						field_column_uom := configs.DB_tbl_mst_item_uom + tglnow.Format("YYYY")
						idrecord_uom_counter := Get_counter(field_column_uom)
						idrecord_uom := tglnow.Format("YY") + tglnow.Format("MM") + strconv.Itoa(idrecord_uom_counter)
						flag_insertuom, msg_insertuom := Exec_SQL(sql_insert, configs.DB_tbl_mst_item_uom, "INSERT",
							idrecord_uom, iditem, iduom,
							default_iduom, convertion,
							admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))
						if flag_insertuom {
							msg = "Succes"
						} else {
							fmt.Println(msg_insertuom)
						}
					}
				} else {
					sql_insert := `
					insert into
					` + configs.DB_tbl_mst_item_uom + ` (
						id_itemuom, iditem, iduom,  
						default_itemuom , conversion_itemuom,
						create_itemuom, createdate_itemuom 
					) values (
						$1, $2, $3,   
						$4, $5, 
						$6, $7 
					)
				`
					field_column_uom := configs.DB_tbl_mst_item_uom + tglnow.Format("YYYY")
					idrecord_uom_counter := Get_counter(field_column_uom)
					idrecord_uom := tglnow.Format("YY") + tglnow.Format("MM") + strconv.Itoa(idrecord_uom_counter)
					flag_insertuom, msg_insertuom := Exec_SQL(sql_insert, configs.DB_tbl_mst_item_uom, "INSERT",
						idrecord_uom, iditem, iduom,
						"N", convertion,
						admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))
					if flag_insertuom {
						msg = "Succes"
					} else {
						fmt.Println(msg_insertuom)
					}
				}
			} else {
				msg = "Duplicate Entry"
			}
		} else {
			msg = "The maximum allowed for this uom is only 3"
		}

	} else {
		if default_iduom == "Y" {
			flag_reset := _reset_itemuom(admin, iditem)
			if flag_reset {
				sql_update := `
					UPDATE 
					` + configs.DB_tbl_mst_item_uom + `  
					SET default_itemuom=$1, conversion_itemuom=$2, 
					update_itemuom=$3, updatedate_itemuom=$4      
					WHERE id_itemuom=$5      
				`

				flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_mst_item_uom, "UPDATE",
					default_iduom, convertion,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

				if flag_update {
					msg = "Succes"
				} else {
					fmt.Println(msg_update)
				}
			}
		} else {
			sql_update := `
				UPDATE 
				` + configs.DB_tbl_mst_item_uom + `  
				SET default_itemuom=$1, conversion_itemuom=$2, 
				update_itemuom=$3, updatedate_itemuom=$4      
				WHERE id_itemuom=$5      
			`

			flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_mst_item_uom, "UPDATE",
				default_iduom, convertion,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
			} else {
				fmt.Println(msg_update)
			}
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Delete_itemuom(idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	render_page := time.Now()
	flag := false
	flag = CheckDB(configs.DB_tbl_mst_item_uom, "id_itemuom", strconv.Itoa(idrecord))

	if flag {
		sql_delete := `
				DELETE FROM
				` + configs.DB_tbl_mst_item_uom + ` 
				WHERE id_itemuom=$1 
			`
		flag_delete, msg_delete := Exec_SQL(sql_delete, configs.DB_tbl_mst_item_uom, "DELETE", idrecord)

		if flag_delete {
			msg = "Succes"
		} else {
			fmt.Println(msg_delete)
		}
	} else {
		msg = "Data Not Found"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}

func _reset_itemuom(admin, iditem string) bool {
	tglnow, _ := goment.New()
	flag := false

	flag_check := CheckDB(configs.DB_tbl_mst_item_uom, "iditem", iditem)
	if flag_check {
		sql_update := `
			UPDATE 
			` + configs.DB_tbl_mst_item_uom + `  
			SET default_itemuom=$1, 
			update_itemuom=$2, updatedate_itemuom=$3     
			WHERE iditem=$4     
		`

		flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_mst_item_uom, "UPDATE",
			"N", admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), iditem)

		if flag_update {
			flag = true
		} else {
			fmt.Println(msg_update + " - _reset_itemuom")
		}
	} else {
		flag = true
	}

	return flag
}
func _Get_merek(idmerek int) string {
	con := db.CreateCon()
	ctx := context.Background()
	nmmerek := ""
	sql_select := `SELECT
			nmmerek    
			FROM ` + database_merek_local + `  
			WHERE idmerek='` + strconv.Itoa(idmerek) + `'       
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&nmmerek); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return nmmerek
}
func _Get_item_uom(iditem string) string {
	con := db.CreateCon()
	ctx := context.Background()
	nmuom := ""
	sql_select := `SELECT
			nmuom    
			FROM ` + configs.DB_tbl_mst_item_uom + `  as A 
			JOIN ` + configs.DB_tbl_mst_uom + `  as B ON B.iduom = A.iduom 
			WHERE iditem='` + iditem + `'     
			AND default_itemuom='Y' LIMIT 1     
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&nmuom); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return nmuom
}
func _total_item_uom(iditem string) int {
	con := db.CreateCon()
	ctx := context.Background()
	total_itemuom := 0
	sql_select := `SELECT
			COUNT(id_itemuom) as total    
			FROM ` + configs.DB_tbl_mst_item_uom + `  
			WHERE iditem='` + iditem + `'     
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&total_itemuom); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return total_itemuom
}
