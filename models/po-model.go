package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/WIGO_MASTER_API/configs"
	"github.com/nikitamirzani323/WIGO_MASTER_API/db"
	"github.com/nikitamirzani323/WIGO_MASTER_API/entities"
	"github.com/nikitamirzani323/WIGO_MASTER_API/helpers"
	"github.com/nleeper/goment"
)

const database_po_local = configs.DB_tbl_trx_po
const database_podetail_local = configs.DB_tbl_trx_po_detail

func Fetch_poHome(search string, page int) (helpers.Responsepaging, error) {
	var obj entities.Model_po
	var arraobj []entities.Model_po
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
	sql_selectcount += "COUNT(idrfq) as totalrfq  "
	sql_selectcount += "FROM " + database_po_local + "  "
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
	sql_select += "A.idpo,  "
	sql_select += "A.idrfq, A.idbranch, B.nmbranch, A.idvendor, C.nmvendor, "
	sql_select += "A.idcurr, A.tipe_docpo, A.statuspo,   "
	sql_select += "A.po_discount, A.po_ppn, A.po_pph,   "
	sql_select += "A.po_totalitem, A.po_subtotal, A.po_grandtotal,   "
	sql_select += "A.createpo, to_char(COALESCE(A.createdatepo,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updatepo, to_char(COALESCE(A.updatedatepo,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_po_local + " as A   "
	sql_select += "JOIN " + configs.DB_tbl_mst_branch + " as B ON B.idbranch = A.idbranch   "
	sql_select += "JOIN " + configs.DB_tbl_mst_vendor + " as C ON C.idvendor = A.idvendor   "
	if search == "" {
		sql_select += "ORDER BY A.createdatepo DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(A.idpo) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdatepo DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpo_db, idrfq_db, idbranch_db, nmbranch_db, idvendor_db, nmvendor_db string
			idcurr_db, tipe_docpo_db, statuspo_db                                 string
			po_discount_db, po_ppn_db, po_pph_db                                  float64
			po_totalitem_db, po_subtotal_db, po_grandtotal_db                     float64
			createpo_db, createdatepo_db, updatepo_db, updatedatepo_db            string
		)

		err = row.Scan(&idpo_db, &idrfq_db, &idbranch_db, &nmbranch_db, &idvendor_db, &nmvendor_db,
			&idcurr_db, &tipe_docpo_db, &statuspo_db,
			&po_discount_db, &po_ppn_db, &po_pph_db,
			&po_totalitem_db, &po_subtotal_db,
			&createpo_db, &createdatepo_db, &updatepo_db, &updatedatepo_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createpo_db != "" {
			create = createpo_db + ", " + createdatepo_db
		}
		if updatepo_db != "" {
			update = updatepo_db + ", " + updatedatepo_db
		}
		switch statuspo_db {
		case "OPEN":
			status_css = configs.STATUS_NEW2
		case "PROCESS":
			status_css = configs.STATUS_RUNNING
		case "COMPLETE":
			status_css = configs.STATUS_COMPLETE
		case "CANCEL":
			status_css = configs.STATUS_CANCEL
		}

		obj.Po_id = idpo_db
		obj.Po_idrfq = idrfq_db
		obj.Po_idbranch = idbranch_db
		obj.Po_idvendor = idvendor_db
		obj.Po_idcurr = idcurr_db
		obj.Po_date = createdatepo_db
		obj.Po_nmbranch = nmbranch_db
		obj.Po_nmvendor = nmvendor_db
		obj.Po_tipedoc = tipe_docpo_db
		obj.Po_discount = po_discount_db
		obj.Po_ppn = po_ppn_db
		obj.Po_ppn = po_pph_db
		obj.Po_totalitem = po_totalitem_db
		obj.Po_subtotal = po_subtotal_db
		obj.Po_grandtotal = po_grandtotal_db
		obj.Po_status = statuspo_db
		obj.Po_status_css = status_css
		obj.Po_create = create
		obj.Po_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_poDetail(idrfq string) (helpers.Response, error) {
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
func Save_po(admin, idrecord, idrfq, sData string, discount, ppn, pph, ppn_total, pph_total, total_item, subtotal, grandtotal float64) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		idbranch_rfq := ""
		idvendor_rfq := ""
		idcurr_rfq := ""
		tipedoc_rfq := ""

		sql_rfq := `SELECT
		 	idbranch, idvendor, idcurr, tipe_documentrfq   
			FROM ` + configs.DB_tbl_trx_rfq + `  
			WHERE idrfq='` + idrfq + `'     
		`
		row_rfq := con.QueryRowContext(ctx, sql_rfq)
		switch e := row_rfq.Scan(&idbranch_rfq, &idvendor_rfq, &idcurr_rfq, &tipedoc_rfq); e {
		case sql.ErrNoRows:
		case nil:
		default:
			helpers.ErrorCheck(e)
		}

		sql_insert := `
				insert into
				` + database_po_local + ` (
					idrfq , idbranch, idvendor, idcurr,  
					po_discount, po_ppn, po_pph, 
					po_ppn_total, po_pph_total, 
					po_totalitem, po_subtotal, po_grandtotal,
					tipe_docpo , statuspo,
					createpo, createdatepo 
				) values (
					$1, $2, $3, $4,     
					$5, $6, $7,
					$8, $9, 
					$10, $11, $12, 
					$13, $14,
					$15, $16 
				)
			`

		field_column := database_po_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := "PO_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		start_date := tglnow.Format("YYYY-MM-DD HH:mm:ss")
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_po_local, "INSERT",
			idrecord, idbranch_rfq, idvendor_rfq, idcurr_rfq,
			discount, ppn, pph,
			ppn_total, pph_total,
			total_item, subtotal, grandtotal,
			tipedoc_rfq, "OPEN",
			admin, start_date)

		if flag_insert {
			msg = "Succes"

			sql_rfqdetail := `SELECT 
				idrfqdetail , idpurchaserequestdetail, idpurchaserequest, 
				iditem , nmitem, descitem, 
				qty , iduom, price 
				FROM ` + configs.DB_tbl_trx_rfq_detail + `  
				WHERE idrfq='` + idrfq + `'  
				AND statusrfqdetail='PROCESS' 
				ORDER BY createdaterfqdetaildetail ASC   `

			row_rfqdetail, err_rfqdetail := con.QueryContext(ctx, sql_rfqdetail)
			helpers.ErrorCheck(err_rfqdetail)
			for row_rfqdetail.Next() {
				var (
					idrfqdetail_db, idpurchaserequestdetail_db, idpurchaserequest_db string
					iditem_db, nmitem_db, descitem_db, iduom_db                      string
					qty_db, price_db                                                 float64
				)

				err_rfqdetail = row_rfqdetail.Scan(&idrfqdetail_db, &idpurchaserequestdetail_db, &idpurchaserequest_db,
					&iditem_db, &nmitem_db, &descitem_db,
					&qty_db, &iduom_db, &price_db)

				_, iddepartement, idemployee, _ := _Get_info_pr(idpurchaserequest_db)

				//admin, idrecord, idpo, idrfqdetail, idrfq, idpurchaserequestdetail, idpurchaserequest,
				// iddepartement, idemployee, iditem, nmitem, descpitem, iduom, status, sData string,
				// qty, price float64
				Save_podetail(admin, "", idrecord, idrfqdetail_db, idrfq,
					idpurchaserequestdetail_db, idpurchaserequest_db, iddepartement, idemployee,
					iditem_db, nmitem_db, descitem_db, iduom_db,
					"OPEN", "New", qty_db, price_db)

			}
			defer row_rfqdetail.Close()

		} else {
			fmt.Println(msg_insert)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_poStatus(admin, idrecord, status string) (helpers.Response, error) {
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
func Save_podetail(admin, idrecord, idpo, idrfqdetail, idrfq, idpurchaserequestdetail, idpurchaserequest, iddepartement, idemployee, iditem, nmitem, descpitem, iduom, status, sData string, qty, price float64) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDBTwoField(database_podetail_local, "idrfq", idrfq, "idpurchaserequestdetail", idpurchaserequestdetail)
		if !flag {
			sql_insert := `
				insert into
				` + database_podetail_local + ` (
					idpodetail, idpo, idrfqdetail, idrfq, 
					idpurchaserequestdetail, idpurchaserequest, iddepartement, idemployee,  
					iditem , nmitem, descitem,  
					qty , iduom, price,  statuspodetail,
					createpodetail, createdaterpodetail 
				) values (
					$1, $2, $3, $4,
					$5, $6, $7, $8, 
					$9, $10, $11,
					$12, $13, $14, $15,    
					$16, $17   
				)
			`
			field_column := database_podetail_local + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			idrecord := "PODETAIL_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_podetail_local, "INSERT",
				idrecord, idpo, idrfqdetail, idrfq,
				idpurchaserequestdetail, idpurchaserequest, iddepartement, idemployee,
				iditem, nmitem, descpitem,
				qty, iduom, price, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
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
func _Get_info_po(idpo string) (string, int) {
	con := db.CreateCon()
	ctx := context.Background()
	status := ""
	total_detail := 0
	sql_select := `SELECT
			statuspo  
			FROM ` + database_po_local + `  
			WHERE idpo='` + idpo + `'     
		`
	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&status); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	sql_selectdetail := `SELECT
			COUNT(idpodetail) AS total 
			FROM ` + database_podetail_local + `  
			WHERE idpo='` + idpo + `'     
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
