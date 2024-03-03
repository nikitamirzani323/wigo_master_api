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

const database_company_local = configs.DB_tbl_mst_company

func Fetch_companyHome(search string, page int) (helpers.Responsercompany, error) {
	var obj entities.Model_company
	var arraobj []entities.Model_company
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	var res helpers.Responsercompany
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
	sql_selectcount += "COUNT(idcompany) as totalcompany  "
	sql_selectcount += "FROM " + database_company_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(idcompany) LIKE '%" + strings.ToLower(search) + "%' "
		sql_selectcount += "OR LOWER(nmcompany) LIKE '%" + strings.ToLower(search) + "%' "
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
	sql_select += "idcompany, "
	sql_select += "startjoincompany, endjoincompany, idcurr, "
	sql_select += "nmcompany, nmowner, emailowner, phone1owner,phone2owner, "
	sql_select += "companyurl_1, companyurl_2, minfee, "
	sql_select += "statuscompany, "
	sql_select += "createcompany, to_char(COALESCE(createdatecompany,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "updatecompany, to_char(COALESCE(updatedatecompany,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_company_local + "   "
	if search == "" {
		sql_select += "ORDER BY createdatecompany DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(idcompany) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "OR LOWER(nmcompany) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY createdatecompany DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcompany_db, startjoincompany_db, endjoincompany_db, idcurr_db                string
			nmcompany_db, nmowner_db, emailowner_db, phone1owner_db, phone2owner_db        string
			companyurl_1_db, companyurl_2_db, statuscompany_db                             string
			minfee_db                                                                      float64
			createcompany_db, createdatecompany_db, updatecompany_db, updatedatecompany_db string
		)

		err = row.Scan(&idcompany_db,
			&startjoincompany_db, &endjoincompany_db, &idcurr_db,
			&nmcompany_db, &nmowner_db, &emailowner_db, &phone1owner_db, &phone2owner_db,
			&companyurl_1_db, &companyurl_2_db, &statuscompany_db, &minfee_db,
			&createcompany_db, &createdatecompany_db, &updatecompany_db, &updatedatecompany_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createcompany_db != "" {
			create = createcompany_db + ", " + createdatecompany_db
		}
		if updatecompany_db != "" {
			update = updatecompany_db + ", " + updatedatecompany_db
		}
		if statuscompany_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Company_id = idcompany_db
		obj.Company_startjoin = startjoincompany_db
		obj.Company_endjoin = endjoincompany_db
		obj.Company_idcurr = idcurr_db
		obj.Company_name = nmcompany_db
		obj.Company_owner = nmowner_db
		obj.Company_email = emailowner_db
		obj.Company_phone1 = phone1owner_db
		obj.Company_phone2 = phone2owner_db
		obj.Company_url1 = companyurl_1_db
		obj.Company_url2 = companyurl_2_db
		obj.Company_minfee = minfee_db
		obj.Company_status = statuscompany_db
		obj.Company_status_css = status_css
		obj.Company_create = create
		obj.Company_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

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
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Listcurr = arraobjcurr
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_company(admin, idrecord, idcurr, nmcompany, nmowner,
	emailowner, phone1, phone2, url1, url2, status, sData string, minfee float64) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_company_local, "idcompany", idrecord)
		if !flag {
			sql_insert := `
				insert into
				` + database_company_local + ` (
					idcompany,  startjoincompany, endjoincompany, idcurr,
					nmcompany , nmowner, emailowner, phone1owner, phone2owner,
					companyurl_1 , companyurl_2, minfee, statuscompany,
					createcompany, createdatecompany 
				) values (
					$1, $2, $3, $4, 
					$5, $6, $7, $8, $9, 
					$10, $11, $12, $13, 
					$14, $15  
				)
			`
			start_join := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_company_local, "INSERT",
				idrecord, start_join, start_join, idcurr,
				idrecord, nmowner, emailowner, phone1, phone2, url1, url2,
				minfee, status,
				nmcompany, start_join)

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
				` + database_company_local + `  
				SET idcurr=$1, 
				nmcompany=$2, nmowner=$3, emailowner=$4, phone1owner=$5, phone2owner=$6,   
				companyurl_1=$7, companyurl_2=$8,  minfee=$9, statuscompany=$10,   
				updatecompany=$11, updatedatecompany=$12    
				WHERE idcompany=$13    
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_company_local, "UPDATE",
			idcurr, nmcompany, nmowner,
			emailowner, phone1, phone2, url1, url2, minfee, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			flag = true
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
