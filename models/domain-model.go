package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/wigo_master_api/configs"
	"bitbucket.org/isbtotogroup/wigo_master_api/db"
	"bitbucket.org/isbtotogroup/wigo_master_api/entities"
	"bitbucket.org/isbtotogroup/wigo_master_api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

const database_domain_local = configs.DB_tbl_mst_domain

func Fetch_domainHome() (helpers.Response, error) {
	var obj entities.Model_domain
	var arraobj []entities.Model_domain
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			iddomain , nmdomain, tipedomain, statusdomain,  
			createdomain, to_char(COALESCE(createdatedomain,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updatedomain, to_char(COALESCE(updatedatedomain,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_domain_local + `  
			ORDER BY createdatedomain DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			iddomain_db                                                                int
			nmdomain_db, tipedomain_db, statusdomain_db                                string
			createdomain_db, createdatedomain_db, updatedomain_db, updatedatedomain_db string
		)

		err = row.Scan(&iddomain_db, &nmdomain_db, &tipedomain_db, &statusdomain_db,
			&createdomain_db, &createdatedomain_db, &updatedomain_db, &updatedatedomain_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createdomain_db != "" {
			create = createdomain_db + ", " + createdatedomain_db
		}
		if updatedomain_db != "" {
			update = updatedomain_db + ", " + updatedatedomain_db
		}
		if statusdomain_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Domain_id = iddomain_db
		obj.Domain_tipe = tipedomain_db
		obj.Domain_name = nmdomain_db
		obj.Domain_status = statusdomain_db
		obj.Domain_status_css = status_css
		obj.Domain_create = create
		obj.Domain_update = update
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
func Fetch_checkdomain(domain, tipe string) (bool, error) {
	ctx := context.Background()
	con := db.CreateCon()
	flag := true
	var result entities.Model_checkdomain
	var nmdomain, tipedomain, statusdomain string
	field_redis := "DOMAINLIST:" + strings.Replace(domain, ":", "_", -1) + "-" + tipe

	_, flagRedis := helpers.GetRedis(field_redis)

	if !flagRedis {
		sql_select := `
			SELECT
			nmdomain, tipedomain, statusdomain  
			FROM ` + database_domain_local + ` 
			WHERE nmdomain = $1 
			AND tipedomain = $2 
			AND statusdomain = 'Y' 
		`

		row := con.QueryRowContext(ctx, sql_select, domain, tipe)
		switch e := row.Scan(&nmdomain, &tipedomain, &statusdomain); e {
		case sql.ErrNoRows:
			return false, errors.New("domain is not registered")
		case nil:
			flag = true
			result.Domain_name = nmdomain
			result.Domain_tipe = tipedomain
			result.Domain_status = statusdomain
			helpers.SetRedis(field_redis, result, 5*time.Hour)
		default:
			return false, errors.New("domain is not registered")
		}
	}

	return flag, nil
}
func Save_domain(admin, tipe, name, status, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_domain_local + ` (
					iddomain , tipedomain, nmdomain, statusdomain,  
					createdomain, createdatedomain 
				) values (
					$1, $2, $3, $4,     
					$5, $6  
				)
			`

		field_column := database_domain_local
		idrecord_counter := Get_counter(field_column)
		idrecord := tglnow.Format("YY") + strconv.Itoa(idrecord_counter)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_domain_local, "INSERT",
			idrecord, tipe, name, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_domain_local + `  
				SET nmdomain=$1, tipedomain=$2, statusdomain=$3, 
				updatedomain=$4, updatedatedomain=$5     
				WHERE iddomain=$6    
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_domain_local, "UPDATE",
			name, tipe, status,
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
