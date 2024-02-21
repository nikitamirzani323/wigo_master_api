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

const database_departement_local = configs.DB_tbl_mst_departement

func Fetch_departementHome(search string, page int) (helpers.Responsepaging, error) {
	var obj entities.Model_departement
	var arraobj []entities.Model_departement
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
	sql_selectcount += "COUNT(iddepartement) as totaldepartement  "
	sql_selectcount += "FROM " + database_departement_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(iddepartement) LIKE '%" + strings.ToLower(search) + "%' "
		sql_selectcount += "OR LOWER(nmdepartement) LIKE '%" + strings.ToLower(search) + "%' "
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
	sql_select += "iddepartement , nmdepartement, statusdepartement, "
	sql_select += "createdepartement, to_char(COALESCE(createdatedepartement,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "updatedepartement, to_char(COALESCE(updatedatedepartement,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_departement_local + "   "
	if search == "" {
		sql_select += "ORDER BY createdatedepartement DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(iddepartement) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "OR LOWER(nmdepartement) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY createdatedepartement DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			iddepartement_db, nmdepartement_db, statusdepartement_db                                       string
			createdepartement_db, createdatedepartement_db, updatedepartement_db, updatedatedepartement_db string
		)

		err = row.Scan(&iddepartement_db, &nmdepartement_db, &statusdepartement_db,
			&createdepartement_db, &createdatedepartement_db, &updatedepartement_db, &updatedatedepartement_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createdepartement_db != "" {
			create = createdepartement_db + ", " + createdatedepartement_db
		}
		if updatedepartement_db != "" {
			update = updatedepartement_db + ", " + updatedatedepartement_db
		}
		if statusdepartement_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Departement_id = iddepartement_db
		obj.Departement_name = nmdepartement_db
		obj.Departement_status = statusdepartement_db
		obj.Departement_status_css = status_css
		obj.Departement_create = create
		obj.Departement_update = update
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
func Save_departement(admin, idrecord, name, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_departement_local, "iddepartement", strings.ToUpper(idrecord))
		if !flag {
			sql_insert := `
				insert into
				` + database_departement_local + ` (
					iddepartement , nmdepartement, statusdepartement, 
					createdepartement, createdatedepartement 
				) values (
					$1, $2, $3,    
					$4, $5  
				)
			`

			flag_insert, msg_insert := Exec_SQL(sql_insert, database_departement_local, "INSERT",
				strings.ToUpper(idrecord), name, status,
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
				` + database_departement_local + `  
				SET nmdepartement=$1, statusdepartement=$2, 
				updatedepartement=$3, updatedatedepartement=$4        
				WHERE iddepartement=$5      
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_departement_local, "UPDATE",
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
