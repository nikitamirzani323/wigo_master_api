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

const database_warehouse_local = configs.DB_tbl_mst_warehouse
const database_warehouse_storage_local = configs.DB_tbl_mst_warehouse_storage
const database_warehouse_storagebin_local = configs.DB_tbl_mst_warehouse_storagebin

func Fetch_warehouseHome(idbranch string) (helpers.Responsewarehouse, error) {
	var obj entities.Model_warehouse
	var arraobj []entities.Model_warehouse
	var objbranch entities.Model_branchshare
	var arraobjbranch []entities.Model_branchshare
	var res helpers.Responsewarehouse
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "A.idwarehouse , A.idbranch, B.nmbranch,  "
	sql_select += "A.nmwarehouse , A.alamatwarehouse, A.phone1warehouse, A.phone2warehouse, A.statuswarehouse,   "
	sql_select += "A.createwarehouse, to_char(COALESCE(A.createdatewarehouse,now()), 'YYYY-MM-DD HH24:MI:SS'),   "
	sql_select += "A.updatewarehouse, to_char(COALESCE(A.updatedatewarehouse,now()), 'YYYY-MM-DD HH24:MI:SS')   "
	sql_select += "FROM " + database_warehouse_local + " AS A "
	sql_select += "JOIN " + configs.DB_tbl_mst_branch + " AS B ON B.idbranch = A.idbranch "
	if idbranch != "" {
		sql_select += "WHERE A.idbranch='" + idbranch + "' "
	}
	sql_select += "ORDER BY A.createdatewarehouse DESC "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idwarehouse_db, idbranch_db, nmbranch_db                                                       string
			nmwarehouse_db, alamatwarehouse_db, phone1warehouse_db, phone2warehouse_db, statuswarehouse_db string
			createwarehouse_db, createdatewarehouse_db, updatewarehouse_db, updatedatewarehouse_db         string
		)

		err = row.Scan(&idwarehouse_db, &idbranch_db, &nmbranch_db,
			&nmwarehouse_db, &alamatwarehouse_db, &phone1warehouse_db, &phone2warehouse_db, &statuswarehouse_db,
			&createwarehouse_db, &createdatewarehouse_db, &updatewarehouse_db, &updatedatewarehouse_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createwarehouse_db != "" {
			create = createwarehouse_db + ", " + createdatewarehouse_db
		}
		if updatewarehouse_db != "" {
			update = updatewarehouse_db + ", " + updatedatewarehouse_db
		}
		if statuswarehouse_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Warehouse_id = idwarehouse_db
		obj.Warehouse_idbranch = idbranch_db
		obj.Warehouse_nmbranch = nmbranch_db
		obj.Warehouse_name = nmwarehouse_db
		obj.Warehouse_alamat = alamatwarehouse_db
		obj.Warehouse_phone1 = phone1warehouse_db
		obj.Warehouse_phone2 = phone2warehouse_db
		obj.Warehouse_status = statuswarehouse_db
		obj.Warehouse_status_css = status_css
		obj.Warehouse_create = create
		obj.Warehouse_update = update
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

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listbranch = arraobjbranch
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_warehouseStorage(idwarehouse string) (helpers.Response, error) {
	var obj entities.Model_warehousestorage
	var arraobj []entities.Model_warehousestorage
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idstorage , nmstorage, statusstorage,  "
	sql_select += "createstorage, to_char(COALESCE(createdatestorage,now()), 'YYYY-MM-DD HH24:MI:SS'),   "
	sql_select += "updatestorage, to_char(COALESCE(updatedatestorage,now()), 'YYYY-MM-DD HH24:MI:SS')   "
	sql_select += "FROM " + database_warehouse_storage_local + " AS A "
	sql_select += "WHERE idwarehouse='" + idwarehouse + "' "
	sql_select += "ORDER BY createdatestorage DESC "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idstorage_db, nmstorage_db, statusstorage_db                                   string
			createstorage_db, createdatestorage_db, updatestorage_db, updatedatestorage_db string
		)

		err = row.Scan(&idstorage_db, &nmstorage_db, &statusstorage_db,
			&createstorage_db, &createdatestorage_db, &updatestorage_db, &updatedatestorage_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createstorage_db != "" {
			create = createstorage_db + ", " + createdatestorage_db
		}
		if updatestorage_db != "" {
			update = updatestorage_db + ", " + updatedatestorage_db
		}
		if statusstorage_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Warehousestorage_id = idstorage_db
		obj.Warehousestorage_name = nmstorage_db
		obj.Warehousestorage_totalbin = _Get_total_storage(idstorage_db)
		obj.Warehousestorage_status = statusstorage_db
		obj.Warehousestorage_status_css = status_css
		obj.Warehousestorage_create = create
		obj.Warehousestorage_update = update
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
func Fetch_warehouseStorageBin(idstorage string) (helpers.Responsestoragebin, error) {
	var obj entities.Model_warehousestoragebin
	var arraobj []entities.Model_warehousestoragebin
	var objuom entities.Model_uomshare
	var arraobjuom []entities.Model_uomshare
	var res helpers.Responsestoragebin
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idbin, nmbin,   "
	sql_select += "iduom, totalcapacity , maxcapacity, statusbin,  "
	sql_select += "createbin, to_char(COALESCE(createdatebin,now()), 'YYYY-MM-DD HH24:MI:SS'),   "
	sql_select += "updatebin, to_char(COALESCE(updatedatebin,now()), 'YYYY-MM-DD HH24:MI:SS')   "
	sql_select += "FROM " + database_warehouse_storagebin_local + " AS A "
	sql_select += "WHERE idstorage='" + idstorage + "' "
	sql_select += "ORDER BY createdatebin DESC "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idbin_db                                                       int
			nmbin_db, iduom_db, statusbin_db                               string
			totalcapacity_db, maxcapacity_db                               float32
			createbin_db, createdatebin_db, updatebin_db, updatedatebin_db string
		)

		err = row.Scan(&idbin_db, &nmbin_db,
			&iduom_db, &totalcapacity_db, &maxcapacity_db, &statusbin_db,
			&createbin_db, &createdatebin_db, &updatebin_db, &updatedatebin_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createbin_db != "" {
			create = createbin_db + ", " + createdatebin_db
		}
		if updatebin_db != "" {
			update = updatebin_db + ", " + updatedatebin_db
		}
		if statusbin_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Warehousestoragebin_id = idbin_db
		obj.Warehousestoragebin_iduom = iduom_db
		obj.Warehousestoragebin_name = nmbin_db
		obj.Warehousestoragebin_maxcapacity = maxcapacity_db
		obj.Warehousestoragebin_totalcapacity = totalcapacity_db
		obj.Warehousestoragebin_status = statusbin_db
		obj.Warehousestoragebin_status_css = status_css
		obj.Warehousestoragebin_create = create
		obj.Warehousestoragebin_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	sql_selectuom := `SELECT 
			iduom, nmuom   
			FROM ` + configs.DB_tbl_mst_uom + ` 
			WHERE statusuom = 'Y' 
			ORDER BY nmuom ASC    
	`
	rowuom, erruom := con.QueryContext(ctx, sql_selectuom)
	helpers.ErrorCheck(erruom)
	for rowuom.Next() {
		var (
			iduom_db, nmuom_db string
		)

		erruom = rowuom.Scan(&iduom_db, &nmuom_db)

		helpers.ErrorCheck(erruom)

		objuom.Uom_id = iduom_db
		objuom.Uom_name = nmuom_db
		arraobjuom = append(arraobjuom, objuom)
		msg = "Success"
	}
	defer rowuom.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listuom = arraobjuom
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_warehouse(admin, idrecord, idbranch, name, alamat, phone1, phone2, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		idrecord_new := strings.ToUpper(idbranch) + "-" + strings.ToUpper(idrecord)
		flag = CheckDB(database_warehouse_local, "idwarehouse", idrecord_new)
		if !flag {
			sql_insert := `
				insert into
				` + database_warehouse_local + ` (
					idwarehouse , idbranch, 
					nmwarehouse , alamatwarehouse, phone1warehouse, phone2warehouse, statuswarehouse,  
					createwarehouse, createdatewarehouse 
				) values (
					$1, $2,    
					$3, $4, $5, $6, $7,    
					$8, $9   
				)
			`
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_warehouse_local, "INSERT",
				idrecord_new, idbranch,
				name, alamat, phone1, phone2, status,
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
				` + database_warehouse_local + `  
				SET nmwarehouse=$1, alamatwarehouse=$2, phone1warehouse=$3, phone2warehouse=$4, statuswarehouse=$5, 
				updatewarehouse=$6, updatedatewarehouse=$7      
				WHERE idwarehouse=$8 
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_warehouse_local, "UPDATE",
			name, alamat, phone1, phone2, status,
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
func Save_warehousestorage(admin, idrecord, idwarehouse, name, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		idrecord_new := strings.ToUpper(idwarehouse) + "-" + strings.ToUpper(idrecord)
		flag = CheckDB(database_warehouse_storage_local, "idstorage", idrecord_new)
		if !flag {
			sql_insert := `
				insert into
				` + database_warehouse_storage_local + ` (
					idstorage , idwarehouse, 
					nmstorage, statusstorage,  
					createstorage, createdatestorage 
				) values (
					$1, $2,    
					$3, $4, 
					$5, $6    
				)
			`

			flag_insert, msg_insert := Exec_SQL(sql_insert, database_warehouse_storage_local, "INSERT",
				idrecord_new, idwarehouse,
				name, status,
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
				` + database_warehouse_storage_local + `  
				SET nmstorage=$1, statusstorage=$2, 
				updatestorage=$3, updatedatestorage=$4       
				WHERE idstorage=$5  
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_warehouse_storage_local, "UPDATE",
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
func Save_warehousestoragebin(admin, idstorage, iduom, name, status, sData string, idrecord int, maxcapacity float32) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
			insert into
			` + database_warehouse_storagebin_local + ` (
				idbin , idstorage, iduom,  
				nmbin, totalcapacity, maxcapacity, statusbin,  
				createbin, createdatebin 
			) values (
				$1, $2, $3,     
				$4, $5, $6, $7,  
				$8, $9    
			)
		`
		field_column := database_warehouse_storagebin_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_warehouse_storage_local, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idstorage, iduom,
			name, 0, maxcapacity, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
			UPDATE 
			` + database_warehouse_storagebin_local + `  
			SET nmbin=$1, maxcapacity=$2, statusbin=$3, 
			updatebin=$4, updatedatebin=$5        
			WHERE idbin=$6   
		`

		flag_update, msg_update := Exec_SQL(sql_update, database_warehouse_storagebin_local, "UPDATE",
			name, maxcapacity, status,
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
func _Get_total_storage(idstorage string) int {
	con := db.CreateCon()
	ctx := context.Background()
	total := 0
	sql_select := `SELECT
			COUNT(idbin)    
			FROM ` + configs.DB_tbl_mst_warehouse_storagebin + `  
			WHERE idstorage='` + idstorage + `'     
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&total); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return total
}
