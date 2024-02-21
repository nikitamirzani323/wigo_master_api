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

const database_employee_local = configs.DB_tbl_mst_employee

func Fetch_employeeHome(search string, page int) (helpers.Responseemployee, error) {
	var obj entities.Model_employee
	var arraobj []entities.Model_employee
	var objdepartement entities.Model_departementshare
	var arraobjdepartement []entities.Model_departementshare
	var res helpers.Responseemployee
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
	sql_selectcount += "COUNT(idemployee) as totalemployee  "
	sql_selectcount += "FROM " + database_employee_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(idemployee) LIKE '%" + strings.ToLower(search) + "%' "
		sql_selectcount += "OR LOWER(nmemployee) LIKE '%" + strings.ToLower(search) + "%' "
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
	sql_select += "A.idemployee, A.iddepartement, B.nmdepartement, A.nmemployee,  "
	sql_select += "A.alamatemployee, A.emailemployee, A.phone1employee, A.phone2employee, A.statusemployee, "
	sql_select += "A.createemployee, to_char(COALESCE(A.createdateemployee,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updateemployee, to_char(COALESCE(A.updatedateemployee,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_employee_local + " as A   "
	sql_select += "JOIN " + configs.DB_tbl_mst_departement + " as B ON B.iddepartement = A.iddepartement   "
	if search == "" {
		sql_select += "ORDER BY A.createdateemployee DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(A.idemployee) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "OR LOWER(A.nmemployee) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdateemployee DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idemployee_db, iddepartement_db, nmdepartement_db, nmemployee_db                             string
			alamatemployee_db, emailemployee_db, phone1employee_db, phone2employee_db, statusemployee_db string
			createemployee_db, createdateemployee_db, updateemployee_db, updatedateemployee_db           string
		)

		err = row.Scan(&idemployee_db, &iddepartement_db, &nmdepartement_db, &nmemployee_db,
			&alamatemployee_db, &emailemployee_db, &phone1employee_db, &phone2employee_db, &statusemployee_db,
			&createemployee_db, &createdateemployee_db, &updateemployee_db, &updatedateemployee_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createemployee_db != "" {
			create = createemployee_db + ", " + createdateemployee_db
		}
		if updateemployee_db != "" {
			update = updateemployee_db + ", " + updatedateemployee_db
		}
		if statusemployee_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Employee_id = idemployee_db
		obj.Employee_iddepartement = iddepartement_db
		obj.Employee_nmdepartement = nmdepartement_db
		obj.Employee_name = nmemployee_db
		obj.Employee_alamat = alamatemployee_db
		obj.Employee_email = emailemployee_db
		obj.Employee_phone1 = phone1employee_db
		obj.Employee_phone2 = phone2employee_db
		obj.Employee_status = statusemployee_db
		obj.Employee_status_css = status_css
		obj.Employee_create = create
		obj.Employee_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

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

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listdepartement = arraobjdepartement
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_employeeShare(iddepartement string) (helpers.Response, error) {
	var obj entities.Model_employeeshare
	var arraobj []entities.Model_employeeshare
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idemployee,nmemployee "
	sql_select += "FROM " + database_employee_local + "    "
	sql_select += "WHERE iddepartement='" + iddepartement + "' "
	sql_select += "AND statusemployee='Y' "
	sql_select += "ORDER BY nmemployee ASC  LIMIT 100 "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idemployee_db, nmemployee_db string
		)

		err = row.Scan(&idemployee_db, &nmemployee_db)

		helpers.ErrorCheck(err)

		obj.Employee_id = idemployee_db
		obj.Employee_name = nmemployee_db
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
func Save_employee(admin, idrecord, iddepart, name, alamat, email, phone1, phone2, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_employee_local + ` (
					idemployee , iddepartement, nmemployee, 
					alamatemployee , emailemployee, phone1employee, phone2employee, statusemployee,
					createemployee, createdateemployee 
				) values (
					$1, $2, $3,    
					$4, $5, $6, $7, $8,   
					$9, $10   
				)
			`
		field_column := database_employee_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_employee_local, "INSERT",
			idrecord, iddepart, name, alamat, email, phone1, phone2, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_employee_local + `  
				SET nmemployee=$1, alamatemployee=$2, 
				emailemployee=$3, phone1employee=$4, phone2employee=$5, statusemployee=$6,
				updateemployee=$7, updatedateemployee=$8         
				WHERE idemployee=$9       
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_employee_local, "UPDATE",
			name, alamat, email, phone1, phone2, status,
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
