package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/db"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nleeper/goment"
)

const database_purchaserequest_local = configs.DB_tbl_trx_purchaserequest
const database_purchaserequestdetail_local = configs.DB_tbl_trx_purchaserequest_detail
const database_prdetail_view_local = configs.DB_view_tbl_pr

func Fetch_purchaserequestHome(search string, page int) (helpers.Responsepurchaserequest, error) {
	var obj entities.Model_purchaserequest
	var arraobj []entities.Model_purchaserequest
	var objbranch entities.Model_branchshare
	var arraobjbranch []entities.Model_branchshare
	var objdepartement entities.Model_departementshare
	var arraobjdepartement []entities.Model_departementshare
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	var res helpers.Responsepurchaserequest
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
	sql_selectcount += "COUNT(idpurchaserequest) as totalpurchase  "
	sql_selectcount += "FROM " + database_purchaserequest_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(idpurchaserequest) LIKE '%" + strings.ToLower(search) + "%' "
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
	sql_select += "A.idpurchaserequest, A.idbranch, B.nmbranch, A.iddepartement, C.nmdepartement, "
	sql_select += "A.idemployee, D.nmemployee, A.idcurr, A.tipe_document, A.periode_document, A.statuspurchaserequest, A.remarkpurchaserequest, A.docexpirepurchaserequest,  "
	sql_select += "A.total_item, A.total_pr , A.total_po,   "
	sql_select += "A.createpurchaserequest, to_char(COALESCE(A.createdatepurchaserequest,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updatepurchaserequest, to_char(COALESCE(A.updatedatepurchaserequest,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_purchaserequest_local + " as A   "
	sql_select += "JOIN " + configs.DB_tbl_mst_branch + " as B ON B.idbranch = A.idbranch   "
	sql_select += "JOIN " + configs.DB_tbl_mst_departement + " as C ON C.iddepartement = A.iddepartement   "
	sql_select += "JOIN " + configs.DB_tbl_mst_employee + " as D ON D.idemployee = A.idemployee   "
	if search == "" {
		sql_select += "ORDER BY A.createdatepurchaserequest DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(A.idpurchaserequest) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdatepurchaserequest DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpurchaserequest_db, idbranch_db, nmbranch_db, iddepartement_db, nmdepartement_db                                                                              string
			idemployee_db, nmemployee_db, idcurr_db, tipe_document_db, periode_document_db, statuspurchaserequest_db, remarkpurchaserequest_db, docexpirepurchaserequest_db string
			total_item_db, total_pr_db, total_po_db                                                                                                                         float64
			createpurchaserequest_db, createdatepurchaserequest_db, updatepurchaserequest_db, updatedatepurchaserequest_db                                                  string
		)

		err = row.Scan(&idpurchaserequest_db, &idbranch_db, &nmbranch_db, &iddepartement_db, &nmdepartement_db,
			&idemployee_db, &nmemployee_db, &idcurr_db, &tipe_document_db, &periode_document_db, &statuspurchaserequest_db, &remarkpurchaserequest_db, &docexpirepurchaserequest_db,
			&total_item_db, &total_pr_db, &total_po_db,
			&createpurchaserequest_db, &createdatepurchaserequest_db, &updatepurchaserequest_db, &updatedatepurchaserequest_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createpurchaserequest_db != "" {
			create = createpurchaserequest_db + ", " + createdatepurchaserequest_db
		}
		if updatepurchaserequest_db != "" {
			update = updatepurchaserequest_db + ", " + updatedatepurchaserequest_db
		}
		switch statuspurchaserequest_db {
		case "OPEN":
			status_css = configs.STATUS_NEW2
		case "PROCESS":
			status_css = configs.STATUS_RUNNING
		case "COMPLETE":
			status_css = configs.STATUS_COMPLETE
		case "CANCEL":
			status_css = configs.STATUS_CANCEL
		}

		obj.Purchaserequest_id = idpurchaserequest_db
		obj.Purchaserequest_date = createdatepurchaserequest_db
		obj.Purchaserequest_idcurr = idcurr_db
		obj.Purchaserequest_idbranch = idbranch_db
		obj.Purchaserequest_nmbranch = nmbranch_db
		obj.Purchaserequest_iddepartement = iddepartement_db
		obj.Purchaserequest_nmdepartement = nmdepartement_db
		obj.Purchaserequest_idemployee = idemployee_db
		obj.Purchaserequest_nmemployee = nmemployee_db
		obj.Purchaserequest_tipedoc = tipe_document_db
		obj.Purchaserequest_periodedoc = periode_document_db
		obj.Purchaserequest_totalitem = total_item_db
		obj.Purchaserequest_totalpr = total_pr_db
		obj.Purchaserequest_totalpo = total_po_db
		obj.Purchaserequest_remark = remarkpurchaserequest_db
		obj.Purchaserequest_docexpire = docexpirepurchaserequest_db
		obj.Purchaserequest_status = statuspurchaserequest_db
		obj.Purchaserequest_status_css = status_css
		obj.Purchaserequest_create = create
		obj.Purchaserequest_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	sql_selectbranch := `SELECT 
			idbranch, nmbranch  
			FROM ` + configs.DB_tbl_mst_branch + ` 
			WHERE statusbranch = 'Y' 
			ORDER BY nmbranch ASC    
	`
	rowbranch, errbranch := con.QueryContext(ctx, sql_selectbranch)
	helpers.ErrorCheck(errbranch)
	for rowbranch.Next() {
		var (
			idbranch_db, nmbranch_db string
		)

		errbranch = rowbranch.Scan(&idbranch_db, &nmbranch_db)

		helpers.ErrorCheck(errbranch)

		objbranch.Branch_id = idbranch_db
		objbranch.Branch_name = nmbranch_db
		arraobjbranch = append(arraobjbranch, objbranch)
		msg = "Success"
	}
	defer rowbranch.Close()

	sql_selectdepartement := `SELECT 
			iddepartement, nmdepartement  
			FROM ` + configs.DB_tbl_mst_departement + ` 
			WHERE statusdepartement = 'Y' 
			ORDER BY nmdepartement ASC    
	`
	rowdepartement, errdepartement := con.QueryContext(ctx, sql_selectdepartement)
	helpers.ErrorCheck(errdepartement)
	for rowdepartement.Next() {
		var (
			iddepartement_db, nmdepartement_db string
		)

		errdepartement = rowdepartement.Scan(&iddepartement_db, &nmdepartement_db)

		helpers.ErrorCheck(errdepartement)

		objdepartement.Departement_id = iddepartement_db
		objdepartement.Departement_name = nmdepartement_db
		arraobjdepartement = append(arraobjdepartement, objdepartement)
		msg = "Success"
	}
	defer rowdepartement.Close()

	sql_selectcurr := `SELECT 
			idcurr 
			FROM ` + configs.DB_tbl_mst_curr + ` 
			ORDER BY idcurr ASC    
	`
	rowcurr, errcurr := con.QueryContext(ctx, sql_selectcurr)
	helpers.ErrorCheck(errcurr)
	for rowcurr.Next() {
		var (
			idcurr_db string
		)

		errcurr = rowcurr.Scan(&idcurr_db)

		helpers.ErrorCheck(errcurr)

		objcurr.Curr_id = idcurr_db
		arraobjcurr = append(arraobjcurr, objcurr)
		msg = "Success"
	}
	defer rowcurr.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listbranch = arraobjbranch
	res.Listdepartement = arraobjdepartement
	res.Listcurr = arraobjcurr
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_purchaserequestDetail(idpurchaserequest string) (helpers.Response, error) {
	var obj entities.Model_purchaserequestdetail
	var arraobj []entities.Model_purchaserequestdetail
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idpurchaserequestdetail, iditem, nmitem, descitem,purpose, "
	sql_select += "qty, iduom, estimateprice, statupurchaserequestdetail,  "
	sql_select += "createpurchaserequestdetail, to_char(COALESCE(createdatepurchaserequestdetail,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "updatepurchaserequestdetail, to_char(COALESCE(updatedatepurchaserequestdetail,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_purchaserequestdetail_local + "   "
	sql_select += "WHERE idpurchaserequest='" + idpurchaserequest + "' "
	sql_select += "ORDER BY createdatepurchaserequestdetail ASC   "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpurchaserequestdetail_db, iditem_db, nmitem_db, descitem_db, purpose_db, iduom_db, statupurchaserequestdetail_db                     string
			qty_db, estimateprice_db                                                                                                               float64
			createpurchaserequestdetail_db, createdatepurchaserequestdetail_db, updatepurchaserequestdetail_db, updatedatepurchaserequestdetail_db string
		)

		err = row.Scan(&idpurchaserequestdetail_db, &iditem_db, &nmitem_db, &descitem_db, &purpose_db,
			&qty_db, &iduom_db, &estimateprice_db, &statupurchaserequestdetail_db,
			&createpurchaserequestdetail_db, &createdatepurchaserequestdetail_db, &updatepurchaserequestdetail_db, &updatedatepurchaserequestdetail_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createpurchaserequestdetail_db != "" {
			create = createpurchaserequestdetail_db + ", " + createdatepurchaserequestdetail_db
		}
		if updatepurchaserequestdetail_db != "" {
			update = updatepurchaserequestdetail_db + ", " + updatedatepurchaserequestdetail_db
		}
		switch statupurchaserequestdetail_db {
		case "OPEN":
			status_css = configs.STATUS_NEW2
		case "PROCESS":
			status_css = configs.STATUS_RUNNING
		case "COMPLETE":
			status_css = configs.STATUS_COMPLETE
		case "CANCEL":
			status_css = configs.STATUS_CANCEL
		}

		obj.Purchaserequestdetail_id = idpurchaserequestdetail_db
		obj.Purchaserequestdetail_iditem = iditem_db
		obj.Purchaserequestdetail_nmitem = nmitem_db
		obj.Purchaserequestdetail_descitem = descitem_db
		obj.Purchaserequestdetail_purpose = purpose_db
		obj.Purchaserequestdetail_iduom = iduom_db
		obj.Purchaserequestdetail_qty = float32(qty_db)
		obj.Purchaserequestdetail_price = float32(estimateprice_db)
		obj.Purchaserequestdetail_status = statupurchaserequestdetail_db
		obj.Purchaserequestdetail_status_css = status_css
		obj.Purchaserequestdetail_create = create
		obj.Purchaserequestdetail_update = update
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
func Fetch_prdetail_view(idbranch, tipedoc string) (helpers.Response, error) {
	var obj entities.Model_prdetail_view
	var arraobj []entities.Model_prdetail_view
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idpurchaserequestdetail, idpurchaserequest,  "
	sql_select += "to_char(COALESCE(createdatepurchaserequest,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "tipe_document, nmbranch, nmdepartement, idemployee, nmemployee,  "
	sql_select += "idcurr, iditem, nmitem, descitem,purpose,  "
	sql_select += "qty, qty_po, iduom, estimateprice, statuspurchaserequest  "
	sql_select += "FROM " + database_prdetail_view_local + "   "
	sql_select += "WHERE tipe_document='" + tipedoc + "' "
	sql_select += "AND idbranch='" + idbranch + "' "
	sql_select += "ORDER BY createdatepurchaserequest ASC   "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpurchaserequestdetail_db, idpurchaserequest_db, createdatepurchaserequest_db               string
			tipe_document_db, nmbranch_db, nmdepartement_db, idemployee_db, nmemployee_db                string
			idcurr_db, iditem_db, nmitem_db, descitem_db, purpose_db, iduom_db, statuspurchaserequest_db string
			qty_db, qty_po_db, estimateprice_db                                                          float64
		)

		err = row.Scan(&idpurchaserequestdetail_db, &idpurchaserequest_db, &createdatepurchaserequest_db,
			&tipe_document_db, &nmbranch_db, &nmdepartement_db, &idemployee_db, &nmemployee_db,
			&idcurr_db, &iditem_db, &nmitem_db, &descitem_db, &purpose_db,
			&qty_db, &qty_po_db, &iduom_db, &estimateprice_db, &statuspurchaserequest_db)
		helpers.ErrorCheck(err)
		status_css := configs.STATUS_CANCEL
		switch statuspurchaserequest_db {
		case "OPEN":
			status_css = configs.STATUS_NEW2
		case "PROCESS":
			status_css = configs.STATUS_RUNNING
		case "COMPLETE":
			status_css = configs.STATUS_COMPLETE
		case "CANCEL":
			status_css = configs.STATUS_CANCEL
		}
		total := qty_db - qty_po_db
		if total > 0 {
			obj.Prdetailview_id = idpurchaserequestdetail_db
			obj.Prdetailview_idpurchaserequest = idpurchaserequest_db
			obj.Prdetailview_date = createdatepurchaserequest_db
			obj.Prdetailview_tipedoc = tipe_document_db
			obj.Prdetailview_nmbranch = nmbranch_db
			obj.Prdetailview_nmdepartement = nmdepartement_db
			obj.Prdetailview_nmemployee = idemployee_db + " - " + nmemployee_db
			obj.Prdetailview_idcurr = idcurr_db
			obj.Prdetailview_iditem = iditem_db
			obj.Prdetailview_nmitem = nmitem_db
			obj.Prdetailview_descitem = descitem_db
			obj.Prdetailview_purpose = purpose_db
			obj.Prdetailview_iduom = iduom_db
			obj.Prdetailview_qty = float32(qty_db)
			obj.Prdetailview_qty_po = float32(qty_po_db)
			obj.Prdetailview_price = float32(estimateprice_db)
			obj.Prdetailview_status = statuspurchaserequest_db
			obj.Prdetailview_status_css = status_css
			arraobj = append(arraobj, obj)
			msg = "Success"
		}

	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_purchaserequest(admin, idrecord, idbranch, iddepartement, idemployee, idcurr, tipedoc, remark, listdetail, sData string, total_item, subtotalpr float32) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_purchaserequest_local + ` (
					idpurchaserequest , idbranch, iddepartement, idemployee, idcurr,  
					tipe_document , periode_document, statuspurchaserequest, remarkpurchaserequest, docexpirepurchaserequest,
					total_item , total_pr,  
					createpurchaserequest, createdatepurchaserequest 
				) values (
					$1, $2, $3, $4, $5,      
					$6, $7, $8, $9, $10, 
					$11, $12, 
					$13, $14  
				)
			`

		field_column := database_purchaserequest_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := "PR_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		periode_doc := tglnow.Format("YYYY") + "-" + tglnow.Format("MM")
		start_date := tglnow.Format("YYYY-MM-DD HH:mm:ss")
		expiredoc := tglnow.Add(2, "days").Format("YYYY-MM-DD HH:mm:ss")
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_purchaserequest_local, "INSERT",
			idrecord, idbranch, iddepartement, idemployee, idcurr,
			tipedoc, periode_doc, "OPEN", remark, expiredoc,
			total_item, subtotalpr,
			admin, start_date)

		if flag_insert {
			msg = "Succes"

			json := []byte(listdetail)
			jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
				detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
				detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
				detail_purpose, _ := jsonparser.GetString(value, "detail_purpose")
				detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
				detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
				detail_price, _ := jsonparser.GetFloat(value, "detail_price")

				Save_purchaserequestdetail(admin, "", idrecord,
					detail_iditem, detail_nmitem, detail_descpitem, detail_purpose, detail_iduom,
					"OPEN", "New", detail_qty, detail_price)
			})
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		_, _, _, totaldetail_db := _Get_info_pr(idrecord)
		if totaldetail_db > 0 {
			sql_delete := `
				DELETE FROM  
				` + database_purchaserequestdetail_local + `   
				WHERE idpurchaserequest=$1  
			`

			flag_delete, msg_delete := Exec_SQL(sql_delete, database_purchaserequestdetail_local, "DELETE", idrecord)

			if flag_delete {
				msg = "Succes"
				//UPDATE
				sql_update := `
					UPDATE 
					` + database_purchaserequest_local + `  
					SET remarkpurchaserequest=$1, total_item=$2, total_pr=$3,   
					updatepurchaserequest=$4, updatedatepurchaserequest=$5        
					WHERE idpurchaserequest=$6      
				`

				flag_update, msg_update := Exec_SQL(sql_update, database_purchaserequest_local, "UPDATE",
					remark, total_item, subtotalpr,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

				if flag_update {
					msg = "Succes"

					json := []byte(listdetail)
					jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
						detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
						detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
						detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
						detail_purpose, _ := jsonparser.GetString(value, "detail_purpose")
						detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
						detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
						detail_price, _ := jsonparser.GetFloat(value, "detail_price")

						Save_purchaserequestdetail(admin, "", idrecord,
							detail_iditem, detail_nmitem, detail_descpitem, detail_purpose, detail_iduom,
							"OPEN", "New", detail_qty, detail_price)
					})
				} else {
					fmt.Println(msg_update)
				}
			} else {
				fmt.Println(msg_delete)
			}
		} else {
			sql_update := `
					UPDATE 
					` + database_purchaserequest_local + `  
					SET remarkpurchaserequest=$1, total_item=$2, total_pr=$3,   
					updatepurchaserequest=$4, updatedatepurchaserequest=$5        
					WHERE idpurchaserequest=$6      
				`

			flag_update, msg_update := Exec_SQL(sql_update, database_purchaserequest_local, "UPDATE",
				remark, total_item, subtotalpr,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"

				json := []byte(listdetail)
				jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
					detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
					detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
					detail_purpose, _ := jsonparser.GetString(value, "detail_purpose")
					detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
					detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
					detail_price, _ := jsonparser.GetFloat(value, "detail_price")

					Save_purchaserequestdetail(admin, "", idrecord,
						detail_iditem, detail_nmitem, detail_descpitem, detail_purpose, detail_iduom,
						"OPEN", "New", detail_qty, detail_price)
				})
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
func Save_purchaserequestStatus(admin, idrecord, status string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	status_db := ""
	total_detail_db := 0

	status_db, _, _, total_detail_db = _Get_info_pr(idrecord)
	if status_db == "OPEN" {
		if total_detail_db > 0 {
			sql_update := `
				UPDATE 
				` + database_purchaserequest_local + `  
				SET statuspurchaserequest=$1, 
				updatepurchaserequest=$2, updatedatepurchaserequest=$3     
				WHERE idpurchaserequest=$4    
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_purchaserequest_local, "UPDATE",
				status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
				//DETAIL
				sql_updatedetail := `
					UPDATE 
					` + database_purchaserequestdetail_local + `  
					SET statupurchaserequestdetail=$1, 
					updatepurchaserequestdetail=$2, updatedatepurchaserequestdetail=$3     
					WHERE idpurchaserequest=$4    
				`

				flag_updatedetail, msg_updatedetail := Exec_SQL(sql_updatedetail, database_purchaserequestdetail_local, "UPDATE",
					status,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

				if flag_updatedetail {
					msg = "Succes"
				} else {
					fmt.Println(msg_updatedetail)
				}
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
func Save_purchaserequestdetail(admin, idrecord, idpurchaserequest, iditem, nmitem, descpitem, purpose, iduom, status, sData string, qty, estimateprice float64) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDBTwoField(database_purchaserequestdetail_local, "idpurchaserequest", idpurchaserequest, "iditem", iditem)
		if !flag {
			sql_insert := `
				insert into
				` + database_purchaserequestdetail_local + ` (
					idpurchaserequestdetail, idpurchaserequest ,  
					iditem , nmitem, descitem,  purpose,
					qty , iduom, estimateprice,  statupurchaserequestdetail,
					createpurchaserequestdetail, createdatepurchaserequestdetail 
				) values (
					$1, $2, 
					$3, $4, $5, $6, 
					$7, $8, $9, $10,    
					$11, $12 
				)
			`
			field_column := database_purchaserequestdetail_local + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			idrecord := "PRDETAIL_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_purchaserequestdetail_local, "INSERT",
				idrecord, idpurchaserequest,
				iditem, nmitem, descpitem, purpose,
				qty, iduom, estimateprice, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_purchaserequestdetail_local + `  
				SET iditem=$1, nmitem=$2, descitem=$3,  purpose=$4,
				qty=$5 , iduom=$6, estimateprice=$7,  statupurchaserequestdetail=$8,
				updatepurchaserequest=$9, updatedatepurchaserequest=$10          
				WHERE idpurchaserequest=$11       
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_purchaserequestdetail_local, "UPDATE",
			iditem, nmitem, descpitem, purpose,
			qty, iduom, estimateprice, status,
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
func _Get_info_pr(idpurchaserequest string) (string, string, string, int) {
	con := db.CreateCon()
	ctx := context.Background()
	iddepartement := ""
	idemployee := ""
	status := ""
	total_detail := 0
	sql_select := `SELECT
			iddepartement,idemployee,statuspurchaserequest  
			FROM ` + database_purchaserequest_local + `  
			WHERE idpurchaserequest='` + idpurchaserequest + `'     
		`
	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&iddepartement, &idemployee, &status); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	sql_selectdetail := `SELECT
			COUNT(idpurchaserequestdetail) AS total 
			FROM ` + database_purchaserequestdetail_local + `  
			WHERE idpurchaserequest='` + idpurchaserequest + `'     
		`
	rowdetail := con.QueryRowContext(ctx, sql_selectdetail)
	switch e := rowdetail.Scan(&total_detail); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return status, iddepartement, idemployee, total_detail
}
func _Get_info_prdetail(idpr, idprdetail string) float64 {
	con := db.CreateCon()
	ctx := context.Background()
	qty_db := 0
	qty_po_db := 0
	qty := 0
	sql_select := `SELECT
			qty, qty_po  
			FROM ` + database_purchaserequestdetail_local + `  
			WHERE idpurchaserequest='` + idpr + `'     
			AND  idpurchaserequestdetail='` + idprdetail + `'     
		`
	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&qty_db, &qty_po_db); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	qty = qty_db - qty_po_db

	return float64(qty)
}
