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

const database_catevendor_local = configs.DB_tbl_mst_catevendor
const database_vendor_local = configs.DB_tbl_mst_vendor

func Fetch_catevendorHome(search string, page int) (helpers.Responsepaging, error) {
	var obj entities.Model_catevendor
	var arraobj []entities.Model_catevendor
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
	sql_selectcount += "COUNT(idcatevendor) as totalcatevendor  "
	sql_selectcount += "FROM " + database_catevendor_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(nmcatevendor) LIKE '%" + strings.ToLower(search) + "%' "
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
	sql_select += "idcatevendor , nmcatevendor, statusvendor, "
	sql_select += "createcatevendor, to_char(COALESCE(createdatecatevendor,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "updatecatevendor, to_char(COALESCE(updatedatecatevendor,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_catevendor_local + "   "
	if search == "" {
		sql_select += "ORDER BY createdatecatevendor DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(nmcatevendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY createdatecatevendor DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcatevendor_db                                                                            int
			nmcatevendor_db, statusvendor_db                                                           string
			createcatevendor_db, createdatecatevendor_db, updatecatevendor_db, updatedatecatevendor_db string
		)

		err = row.Scan(&idcatevendor_db, &nmcatevendor_db, &statusvendor_db,
			&createcatevendor_db, &createdatecatevendor_db, &updatecatevendor_db, &updatedatecatevendor_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createcatevendor_db != "" {
			create = createcatevendor_db + ", " + createdatecatevendor_db
		}
		if updatecatevendor_db != "" {
			update = updatecatevendor_db + ", " + updatedatecatevendor_db
		}
		if statusvendor_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Catevendor_id = idcatevendor_db
		obj.Catevendor_name = nmcatevendor_db
		obj.Catevendor_status = statusvendor_db
		obj.Catevendor_status_css = status_css
		obj.Catevendor_create = create
		obj.Catevendor_update = update
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
func Fetch_vendorHome(search string, page int) (helpers.Responsevendor, error) {
	var obj entities.Model_vendor
	var arraobj []entities.Model_vendor
	var objcatevendor entities.Model_catevendorshare
	var arraobjcatevendor []entities.Model_catevendorshare
	var res helpers.Responsevendor
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
	sql_selectcount += "COUNT(idvendor) as totalvendor  "
	sql_selectcount += "FROM " + database_vendor_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(idvendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_selectcount += "OR LOWER(nmvendor) LIKE '%" + strings.ToLower(search) + "%' "
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
	sql_select += "A.idvendor ,A.idcatevendor,B.nmcatevendor ,A.nmvendor, A.picvendor, "
	sql_select += "A.alamatvendor, A.emailvendor, A.phone1vendor, A.phone2vendor, A.statusvendor,  "
	sql_select += "A.createvendor, to_char(COALESCE(A.createdatevendor,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updatevendor, to_char(COALESCE(A.updatedatevendor,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_vendor_local + " as A   "
	sql_select += "JOIN " + database_catevendor_local + " as B ON B.idcatevendor = A.idcatevendor   "
	if search == "" {
		sql_select += "ORDER BY A.createdatevendor DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(A.idvendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "OR LOWER(A.nmvendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdatevendor DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcatevendor_db                                                                    int
			idvendor_db, nmcatevendor_db, nmvendor_db, picvendor_db                            string
			alamatvendor_db, emailvendor_db, phone1vendor_db, phone2vendor_db, statusvendor_db string
			createvendor_db, createdatevendor_db, updatevendor_db, updatedatevendor_db         string
		)

		err = row.Scan(&idvendor_db, &idcatevendor_db, &nmcatevendor_db, &nmvendor_db, &picvendor_db, &alamatvendor_db,
			&emailvendor_db, &phone1vendor_db, &phone2vendor_db, &statusvendor_db,
			&createvendor_db, &createdatevendor_db, &updatevendor_db, &updatedatevendor_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createvendor_db != "" {
			create = createvendor_db + ", " + createdatevendor_db
		}
		if updatevendor_db != "" {
			update = updatevendor_db + ", " + updatedatevendor_db
		}
		if statusvendor_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Vendor_id = idvendor_db
		obj.Vendor_idcatevendor = idcatevendor_db
		obj.Vendor_nmcatevendor = nmcatevendor_db
		obj.Vendor_name = nmvendor_db
		obj.Vendor_pic = picvendor_db
		obj.Vendor_alamat = alamatvendor_db
		obj.Vendor_email = emailvendor_db
		obj.Vendor_phone1 = phone1vendor_db
		obj.Vendor_phone2 = phone2vendor_db
		obj.Vendor_status = statusvendor_db
		obj.Vendor_status_css = status_css
		obj.Vendor_create = create
		obj.Vendor_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	sql_selectcatevendor := `SELECT 
			idcatevendor, nmcatevendor  
			FROM ` + database_catevendor_local + ` 
			WHERE statusvendor = 'Y' 
			ORDER BY nmcatevendor ASC    
	`
	rowcatevendor, errcatevendor := con.QueryContext(ctx, sql_selectcatevendor)
	helpers.ErrorCheck(errcatevendor)
	for rowcatevendor.Next() {
		var (
			idcatevendor_db int
			nmcatevendor_db string
		)

		errcatevendor = rowcatevendor.Scan(&idcatevendor_db, &nmcatevendor_db)

		helpers.ErrorCheck(errcatevendor)

		objcatevendor.Catevendor_id = idcatevendor_db
		objcatevendor.Catevendor_name = nmcatevendor_db
		arraobjcatevendor = append(arraobjcatevendor, objcatevendor)
		msg = "Success"
	}
	defer rowcatevendor.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listcatevendor = arraobjcatevendor
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_vendorShare(search string) (helpers.Response, error) {
	var obj entities.Model_vendorshare
	var arraobj []entities.Model_vendorshare
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "A.idvendor ,B.nmcatevendor ,A.nmvendor "
	sql_select += "FROM " + database_vendor_local + " as A   "
	sql_select += "JOIN " + database_catevendor_local + " as B ON B.idcatevendor = A.idcatevendor   "
	sql_select += "WHERE A.statusvendor='Y' "
	if search == "" {
		sql_select += "ORDER BY A.createdatevendor DESC   LIMIT " + strconv.Itoa(configs.PAGING_PAGE)
	} else {
		sql_select += "AND LOWER(A.nmvendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdatevendor DESC   LIMIT " + strconv.Itoa(configs.PAGING_PAGE)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idvendor_db, nmcatevendor_db, nmvendor_db string
		)

		err = row.Scan(&idvendor_db, &nmcatevendor_db, &nmvendor_db)

		helpers.ErrorCheck(err)

		obj.Vendor_id = idvendor_db
		obj.Vendor_nmcatevendor = nmcatevendor_db
		obj.Vendor_name = nmvendor_db
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
func Save_catevendor(admin, name, status, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_catevendor_local + ` (
					idcatevendor , nmcatevendor, statusvendor, 
					createcatevendor, createdatecatevendor 
				) values (
					$1, $2, $3,    
					$4, $5 
				)
			`
		field_column := database_catevendor_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := tglnow.Format("YY") + strconv.Itoa(idrecord_counter)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_vendor_local, "INSERT",
			idrecord, name, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_catevendor_local + `  
				SET nmcatevendor=$1, statusvendor=$2, 
				updatecatevendor=$3, updatedatecatevendor=$4        
				WHERE idcatevendor=$5     
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_catevendor_local, "UPDATE",
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
func Save_vendor(admin, idrecord, name, pic, alamat, email, phone1, phone2, status, sData string, idcatevendor int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_vendor_local + ` (
					idvendor , idcatevendor, nmvendor, picvendor, 
					alamatvendor,emailvendor, phone1vendor, phone2vendor, statusvendor,
					createvendor, createdatevendor 
				) values (
					$1, $2, $3, $4,     
					$5, $6, $7, $8, $9,  
					$10, $11  
				)
			`
		field_column := database_vendor_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := "VENDOR_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_vendor_local, "INSERT",
			idrecord, idcatevendor, name, pic,
			alamat, email, phone1, phone2, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_vendor_local + `  
				SET idcatevendor=$1, nmvendor=$2, picvendor=$3, alamatvendor=$4,  
				emailvendor=$5, phone1vendor=$6, phone2vendor=$7 ,statusvendor=$8, 
				updatevendor=$9, updatedatevendor=$10         
				WHERE idvendor=$11     
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_vendor_local, "UPDATE",
			idcatevendor, name, pic, alamat, email, phone1, phone2, status,
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
