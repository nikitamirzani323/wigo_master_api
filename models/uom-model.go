package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/WIGO_MASTER_API/configs"
	"github.com/nikitamirzani323/WIGO_MASTER_API/db"
	"github.com/nikitamirzani323/WIGO_MASTER_API/entities"
	"github.com/nikitamirzani323/WIGO_MASTER_API/helpers"
	"github.com/nleeper/goment"
)

const database_uom_local = configs.DB_tbl_mst_uom

func Fetch_uomHome() (helpers.Response, error) {
	var obj entities.Model_uom
	var arraobj []entities.Model_uom
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			iduom , nmuom, statusuom, 
			createuom, to_char(COALESCE(createdateuom,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updateuom, to_char(COALESCE(updatedateuom,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_uom_local + `  
			ORDER BY createdateuom DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			iduom_db, nmuom_db, statusuom_db                               string
			createuom_db, createdateuom_db, updateuom_db, updatedateuom_db string
		)

		err = row.Scan(&iduom_db, &nmuom_db, &statusuom_db,
			&createuom_db, &createdateuom_db, &updateuom_db, &updatedateuom_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createuom_db != "" {
			create = createuom_db + ", " + createdateuom_db
		}
		if updateuom_db != "" {
			update = updateuom_db + ", " + updatedateuom_db
		}
		if statusuom_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Uom_id = iduom_db
		obj.Uom_name = nmuom_db
		obj.Uom_status = statusuom_db
		obj.Uom_status_css = status_css
		obj.Uom_create = create
		obj.Uom_update = update
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
func Fetch_uomShare() (helpers.Response, error) {
	var obj entities.Model_uomshare
	var arraobj []entities.Model_uomshare
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			iduom , nmuom 
			FROM ` + database_uom_local + `  
			WHERE statusuom='Y' 
			ORDER BY createdateuom DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			iduom_db, nmuom_db string
		)

		err = row.Scan(&iduom_db, &nmuom_db)

		helpers.ErrorCheck(err)

		obj.Uom_id = iduom_db
		obj.Uom_name = nmuom_db
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
func Save_uom(admin, idrecord, name, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_uom_local, "iduom", strings.ToUpper(idrecord))
		if !flag {
			sql_insert := `
				insert into
				` + database_uom_local + ` (
					iduom , nmuom, statusuom, 
					createuom, createdateuom 
				) values (
					$1, $2, $3,     
					$4, $5   
				)
			`
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_uom_local, "INSERT",
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
				` + database_uom_local + `  
				SET nmuom=$1, statusuom=$2, 
				updateuom=$3, updatedateuom=$4      
				WHERE iduom=$5    
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_uom_local, "UPDATE",
			name, status,
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
