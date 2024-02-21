package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

const database_rfq_local = configs.DB_tbl_trx_rfq
const database_rfqdetail_local = configs.DB_tbl_trx_rfq_detail

func Fetch_rfqHome(search string, page int) (helpers.Responserfq, error) {
	var obj entities.Model_rfq
	var arraobj []entities.Model_rfq
	var objbranch entities.Model_branchshare
	var arraobjbranch []entities.Model_branchshare
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	var res helpers.Responserfq
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
	sql_selectcount += "COUNT(idrfq) as totalrfq  "
	sql_selectcount += "FROM " + database_rfq_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(idrfq) LIKE '%" + strings.ToLower(search) + "%' "
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
	sql_select += "A.idrfq, A.idbranch, B.nmbranch, A.idvendor, C.nmvendor, "
	sql_select += "A.idcurr, A.tipe_documentrfq, A.statusrfq,   "
	sql_select += "A.rfq_total_item, A.rfq_total_rfq,    "
	sql_select += "A.createrfq, to_char(COALESCE(A.createdaterfq,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updaterfq, to_char(COALESCE(A.updatedaterfq,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_rfq_local + " as A   "
	sql_select += "JOIN " + configs.DB_tbl_mst_branch + " as B ON B.idbranch = A.idbranch   "
	sql_select += "JOIN " + configs.DB_tbl_mst_vendor + " as C ON C.idvendor = A.idvendor   "
	if search == "" {
		sql_select += "ORDER BY A.createdaterfq DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(A.idrfq) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdaterfq DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idrfq_db, idbranch_db, nmbranch_db, idvendor_db, nmvendor_db   string
			idcurr_db, tipe_documentrfq_db, statusrfq_db                   string
			rfq_total_item_db, rfq_total_rfq_db                            float64
			createrfq_db, createdaterfq_db, updaterfq_db, updatedaterfq_db string
		)

		err = row.Scan(&idrfq_db, &idbranch_db, &nmbranch_db, &idvendor_db, &nmvendor_db,
			&idcurr_db, &tipe_documentrfq_db, &statusrfq_db,
			&rfq_total_item_db, &rfq_total_rfq_db,
			&createrfq_db, &createdaterfq_db, &updaterfq_db, &updatedaterfq_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createrfq_db != "" {
			create = createrfq_db + ", " + createdaterfq_db
		}
		if updaterfq_db != "" {
			update = updaterfq_db + ", " + updatedaterfq_db
		}
		switch statusrfq_db {
		case "OPEN":
			status_css = configs.STATUS_NEW2
		case "PROCESS":
			status_css = configs.STATUS_RUNNING
		case "COMPLETE":
			status_css = configs.STATUS_COMPLETE
		case "CANCEL":
			status_css = configs.STATUS_CANCEL
		}

		obj.Rfq_id = idrfq_db
		obj.Rfq_date = createdaterfq_db
		obj.Rfq_idbranch = idbranch_db
		obj.Rfq_idvendor = idvendor_db
		obj.Rfq_idcurr = idcurr_db
		obj.Rfq_nmbranch = nmbranch_db
		obj.Rfq_nmvendor = nmvendor_db
		obj.Rfq_tipedoc = tipe_documentrfq_db
		obj.Rfq_totalitem = rfq_total_item_db
		obj.Rfq_totalrfq = rfq_total_rfq_db
		obj.Rfq_status = statusrfq_db
		obj.Rfq_status_css = status_css
		obj.Rfq_create = create
		obj.Rfq_update = update
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
	res.Listcurr = arraobjcurr
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_rfqDetail(idrfq string) (helpers.Response, error) {
	var obj entities.Model_rfqdetail
	var arraobj []entities.Model_rfqdetail
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "A.idrfqdetail, A.idpurchaserequestdetail, A.idpurchaserequest, "
	sql_select += "C.nmdepartement, B.idemployee, D.nmemployee,  "
	sql_select += "A.iditem, A.nmitem, A.descitem, "
	sql_select += "A.qty, A.iduom, A.price, A.statusrfqdetail,  "
	sql_select += "A.createrfqdetail, to_char(COALESCE(A.createdaterfqdetaildetail,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updaterfqdetaildetail, to_char(COALESCE(A.updatedaterfqdetail,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_rfqdetail_local + " as A   "
	sql_select += "JOIN " + configs.DB_tbl_trx_purchaserequest + " as B ON B.idpurchaserequest = A.idpurchaserequest   "
	sql_select += "JOIN " + configs.DB_tbl_mst_departement + " as C ON C.iddepartement = B.iddepartement   "
	sql_select += "JOIN " + configs.DB_tbl_mst_employee + " as D ON D.idemployee = B.idemployee   "
	sql_select += "WHERE A.idrfq='" + idrfq + "' "
	sql_select += "ORDER BY A.createdaterfqdetaildetail ASC   "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idrfqdetail_db, idpurchaserequestdetail_db, idpurchaserequest_db                                   string
			nmdepartement_db, idemployee_db, nmemployee_db                                                     string
			iditem_db, nmitem_db, descitem_db, iduom_db, statusrfqdetail_db                                    string
			qty_db, price_db                                                                                   float64
			createrfqdetail_db, createdaterfqdetaildetail_db, updaterfqdetaildetail_db, updatedaterfqdetail_db string
		)

		err = row.Scan(&idrfqdetail_db, &idpurchaserequestdetail_db, &idpurchaserequest_db,
			&nmdepartement_db, &idemployee_db, &nmemployee_db,
			&iditem_db, &nmitem_db, &descitem_db,
			&qty_db, &iduom_db, &price_db, &statusrfqdetail_db,
			&createrfqdetail_db, &createdaterfqdetaildetail_db, &updaterfqdetaildetail_db, &updatedaterfqdetail_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createrfqdetail_db != "" {
			create = createrfqdetail_db + ", " + createdaterfqdetaildetail_db
		}
		if updaterfqdetaildetail_db != "" {
			update = updaterfqdetaildetail_db + ", " + updatedaterfqdetail_db
		}
		switch statusrfqdetail_db {
		case "OPEN":
			status_css = configs.STATUS_NEW2
		case "PROCESS":
			status_css = configs.STATUS_RUNNING
		case "COMPLETE":
			status_css = configs.STATUS_COMPLETE
		case "CANCEL":
			status_css = configs.STATUS_CANCEL
		}

		obj.Rfqdetail_id = idrfqdetail_db
		obj.Rfqdetail_idpurchaserequestdetail = idpurchaserequestdetail_db
		obj.Rfqdetail_idpurchaserequest = idpurchaserequest_db
		obj.Rfqdetail_nmdepartement = nmdepartement_db
		obj.Rfqdetail_nmemployee = idemployee_db + " - " + nmemployee_db
		obj.Rfqdetail_iditem = iditem_db
		obj.Rfqdetail_nmitem = nmitem_db
		obj.Rfqdetail_descitem = descitem_db
		obj.Rfqdetail_iduom = iduom_db
		obj.Rfqdetail_qty = float64(qty_db)
		obj.Rfqdetail_price = float64(price_db)
		obj.Rfqdetail_status = statusrfqdetail_db
		obj.Rfqdetail_status_css = status_css
		obj.Rfqdetail_create = create
		obj.Rfqdetail_update = update
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
func Save_rfq(admin, idrecord, idbranch, idvendor, idcurr, tipedoc, listdetail, sData string, total_item, subtotalpr float32) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_rfq_local + ` (
					idrfq , idbranch, idvendor, idcurr,  
					tipe_documentrfq , statusrfq, rfq_total_item, rfq_total_rfq,
					createrfq, createdaterfq 
				) values (
					$1, $2, $3, $4,     
					$5, $6, $7, $8, 
					$9, $10  
				)
			`

		field_column := database_rfq_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := "RFQ_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		start_date := tglnow.Format("YYYY-MM-DD HH:mm:ss")
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_rfq_local, "INSERT",
			idrecord, idbranch, idvendor, idcurr,
			tipedoc, "OPEN", total_item, subtotalpr,
			admin, start_date)

		if flag_insert {
			msg = "Succes"

			json := []byte(listdetail)
			jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				detail_id, _ := jsonparser.GetString(value, "detail_id")
				detail_document, _ := jsonparser.GetString(value, "detail_document")
				detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
				detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
				detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
				detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
				detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
				detail_price, _ := jsonparser.GetFloat(value, "detail_price")

				//admin, idrecord, idrfq, idpurchaserequestdetail, idpurchaserequest, iditem, nmitem, descpitem, iduom, status, sData string, qty, price float64
				Save_rfqdetail(admin, "", idrecord,
					detail_id, detail_document,
					detail_iditem, detail_nmitem, detail_descpitem, detail_iduom,
					"OPEN", "New", detail_qty, detail_price)
			})
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		_, totaldetail_db := _Get_info_rfq(idrecord)
		log.Println("total : ", totaldetail_db)
		log.Println("data : ", listdetail)
		if totaldetail_db > 0 {
			sql_delete := `
				DELETE FROM  
				` + database_rfqdetail_local + `   
				WHERE idrfq=$1  
			`

			flag_delete, msg_delete := Exec_SQL(sql_delete, database_rfqdetail_local, "DELETE", idrecord)

			if flag_delete {
				msg = "Succes"
				//UPDATE
				sql_update := `
					UPDATE 
					` + database_rfq_local + `  
					SET rfq_total_item=$1, rfq_total_rfq=$2,   
					updaterfq=$3, updatedaterfq=$4         
					WHERE idrfq=$5       
				`

				flag_update, msg_update := Exec_SQL(sql_update, database_rfq_local, "UPDATE",
					total_item, subtotalpr,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

				if flag_update {
					msg = "Succes"

					json := []byte(listdetail)
					jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
						detail_id, _ := jsonparser.GetString(value, "detail_id")
						detail_document, _ := jsonparser.GetString(value, "detail_document")
						detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
						detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
						detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
						detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
						detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
						detail_price, _ := jsonparser.GetFloat(value, "detail_price")

						//admin, idrecord, idrfq, idpurchaserequestdetail, idpurchaserequest, iditem, nmitem, descpitem, iduom, status, sData string, qty, price float64
						Save_rfqdetail(admin, "", idrecord,
							detail_id, detail_document,
							detail_iditem, detail_nmitem, detail_descpitem, detail_iduom,
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
					` + database_rfq_local + `  
					SET rfq_total_item=$1, rfq_total_rfq=$2,   
					updaterfq=$3, updatedaterfq=$4         
					WHERE idrfq=$5       
				`

			flag_update, msg_update := Exec_SQL(sql_update, database_purchaserequest_local, "UPDATE",
				total_item, subtotalpr,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"

				json := []byte(listdetail)
				jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					detail_id, _ := jsonparser.GetString(value, "detail_id")
					detail_document, _ := jsonparser.GetString(value, "detail_document")
					detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
					detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
					detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
					detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
					detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
					detail_price, _ := jsonparser.GetFloat(value, "detail_price")

					//admin, idrecord, idrfq, idpurchaserequestdetail, idpurchaserequest, iditem, nmitem, descpitem, iduom, status, sData string, qty, price float64
					Save_rfqdetail(admin, "", idrecord,
						detail_id, detail_document,
						detail_iditem, detail_nmitem, detail_descpitem, detail_iduom,
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
func Save_rfqStatus(admin, idrecord, status string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	status_db := ""
	total_detail_db := 0

	status_db, total_detail_db = _Get_info_rfq(idrecord)
	if status_db == "OPEN" {
		if total_detail_db > 0 {
			sql_update := `
				UPDATE 
				` + database_rfq_local + `  
				SET statusrfq=$1, 
				updaterfq=$2, updatedaterfq=$3     
				WHERE idrfq=$4    
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_rfq_local, "UPDATE",
				status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
				//DETAIL
				sql_updatedetail := `
					UPDATE 
					` + database_rfqdetail_local + `  
					SET statusrfqdetail=$1, 
					updaterfqdetaildetail=$2, updatedaterfqdetail=$3     
					WHERE idrfq=$4    
				`

				flag_updatedetail, msg_updatedetail := Exec_SQL(sql_updatedetail, database_rfqdetail_local, "UPDATE",
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
	} else if status_db == "PROCESS" {
		if status == "CANCEL" {
			sql_update := `
				UPDATE 
				` + database_rfq_local + `  
				SET statusrfq=$1, 
				updaterfq=$2, updatedaterfq=$3     
				WHERE idrfq=$4    
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_rfq_local, "UPDATE",
				"CANCEL",
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
				//DETAIL
				sql_updatedetail := `
					UPDATE 
					` + database_rfqdetail_local + `  
					SET statusrfqdetail=$1, 
					updaterfqdetaildetail=$2, updatedaterfqdetail=$3     
					WHERE idrfq=$4    
				`

				flag_updatedetail, msg_updatedetail := Exec_SQL(sql_updatedetail, database_rfqdetail_local, "UPDATE",
					"CANCEL",
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
func Save_rfqdetail(admin, idrecord, idrfq, idpurchaserequestdetail, idpurchaserequest, iditem, nmitem, descpitem, iduom, status, sData string, qty, price float64) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDBTwoField(database_rfqdetail_local, "idrfq", idrfq, "idpurchaserequestdetail", idpurchaserequestdetail)
		if !flag {
			qty_total := _Get_info_prdetail(idpurchaserequest, idpurchaserequestdetail)
			if qty < qty_total {
				sql_insert := `
					insert into
					` + database_rfqdetail_local + ` (
						idrfqdetail, idrfq, 
						idpurchaserequestdetail, idpurchaserequest,  
						iditem , nmitem, descitem,  
						qty , iduom, price,  statusrfqdetail,
						createrfqdetail, createdaterfqdetaildetail 
					) values (
						$1, $2, 
						$3, $4, 
						$5, $6, $7, 
						$8, $9, $10, $11,    
						$12, $13  
					)
				`
				field_column := database_rfqdetail_local + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_column)
				idrecord := "RFQDETAIL_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
				flag_insert, msg_insert := Exec_SQL(sql_insert, database_rfqdetail_local, "INSERT",
					idrecord, idrfq,
					idpurchaserequestdetail, idpurchaserequest,
					iditem, nmitem, descpitem,
					qty, iduom, price, status,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

				if flag_insert {
					msg = "Succes"
				} else {
					fmt.Println(msg_insert)
				}
			} else {
				msg = "Qty exceeds Purchase Qty"
			}
		} else {
			msg = "Duplicate Entry"
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func _Get_info_rfq(idrfq string) (string, int) {
	con := db.CreateCon()
	ctx := context.Background()
	status := ""
	total_detail := 0
	sql_select := `SELECT
			statusrfq  
			FROM ` + database_rfq_local + `  
			WHERE idrfq='` + idrfq + `'     
		`
	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&status); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	sql_selectdetail := `SELECT
			COUNT(idrfqdetail) AS total 
			FROM ` + database_rfqdetail_local + `  
			WHERE idrfq='` + idrfq + `'     
		`
	rowdetail := con.QueryRowContext(ctx, sql_selectdetail)
	switch e := rowdetail.Scan(&total_detail); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return status, total_detail
}
