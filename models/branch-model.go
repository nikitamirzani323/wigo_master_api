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

const database_branch_local = configs.DB_tbl_mst_branch

func Fetch_branchHome() (helpers.Response, error) {
	var obj entities.Model_branch
	var arraobj []entities.Model_branch
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idbranch , nmbranch, statusbranch,  
			createbranch, to_char(COALESCE(createdatebranch,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updatebranch, to_char(COALESCE(updatedatebranch,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_branch_local + `  
			ORDER BY createdatebranch DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idbranch_db, nmbranch_db, statusbranch_db                                  string
			createbranch_db, createdatebranch_db, updatebranch_db, updatedatebranch_db string
		)

		err = row.Scan(&idbranch_db, &nmbranch_db, &statusbranch_db,
			&createbranch_db, &createdatebranch_db, &updatebranch_db, &updatedatebranch_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createbranch_db != "" {
			create = createbranch_db + ", " + createdatebranch_db
		}
		if updatebranch_db != "" {
			update = updatebranch_db + ", " + updatedatebranch_db
		}
		if statusbranch_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Branch_id = idbranch_db
		obj.Branch_name = nmbranch_db
		obj.Branch_status = statusbranch_db
		obj.Branch_status_css = status_css
		obj.Branch_create = create
		obj.Branch_update = update
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
func Save_branch(admin, idrecord, name, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_branch_local, "idbranch", strings.ToUpper(idrecord))
		if !flag {
			sql_insert := `
				insert into
				` + database_branch_local + ` (
					idbranch , nmbranch, statusbranch, 
					createbranch, createdatebranch 
				) values (
					$1, $2, $3,     
					$4, $5   
				)
			`
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_branch_local, "INSERT",
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
				` + database_branch_local + `  
				SET nmbranch=$1, statusbranch=$2, 
				updatebranch=$3, updatedatebranch=$4     
				WHERE idbranch=$5     
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_branch_local, "UPDATE",
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
