package dao

import (
	"database/sql"
	"husen.co.id/model"
)

func InsertGeoTree(db *sql.DB, ncID int64, userParam *model.DataUserLevel) (err error) {
	//funcName := "InsertDataGroup"
	query :=
		" INSERT INTO  geo_tree " +
			"  (nexchief_account_id, code, name," +
			"  level, parent_id ) " +
			" VALUES " +
			"($1, $2, $3, " +
			" $4, $5) "
	param := []interface{}{
		ncID, userParam.Code, userParam.Name,
		userParam.Level, userParam.ParentID}

	_, err = db.Exec(query, param...)

	return
}

func UpdateGeoTree(db *sql.DB, ncID int64, userParam *model.DataUserLevel) (err error) {
	//funcName := "InsertDataGroup"
	query :=
		" UPDATE  geo_tree SET " +
			"nexchief_account_id = $1, code = $2, name = $3," +
			"level = $4, parent_id = $5 " +
			"WHERE id = $6 "
	param := []interface{}{
		ncID, userParam.Code, userParam.Name,
		userParam.Level, userParam.ParentID, userParam.ID}

	_, err = db.Exec(query, param...)

	return
}

func GetID(db *sql.DB, ncID int64, data *model.DataUserLevel) (err error) {
	query := "SELECT id FROM geo_tree WHERE nexchief_account_id = $1 AND code = $2 " +
		"AND level = $3 AND parent_id = $4 "
	param := []interface{}{ncID, data.Code, data.Level, data.ParentID}
	errs := db.QueryRow(query, param...).Scan(data.ID)
	if errs != nil && errs != sql.ErrNoRows {
		err = errs
	}
	return
}
func GetParent(db *sql.DB, ncID int64, data *model.DataUserLevel, listData []interface{}) (err error) {

	query, param := GetQueryParent("geo_tree", "parent", "code", nil, listData, 0)

	rows, err := db.Query(query, param...)
	if err != nil {
		return
	}
	if rows != nil {
		defer func() {
			err = rows.Close()
			if err != nil {
				return
			}
		}()

		for rows.Next() {
			var key string
			var id sql.NullInt64

			err = rows.Scan(&key, &id)
			if err != nil {
				return
			}
			switch key {
			case "parent":
				data.ParentID = id.Int64
			}
		}
	}
	return err
}

