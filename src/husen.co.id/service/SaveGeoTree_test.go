package service
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//	"github.com/stretchr/testify/assert"
//	"github.com/tkanos/gonfig"
//	"rohmat.co.id/dao"
//	"rohmat.co.id/model"
//	"testing"
//)
//
//
//func TestInsertUserLevel1(t *testing.T) {
//	var (
//		temp            model.DataUserLevel
//		nexchiefAccount model.model.NexchiefAccount
//	)
//	err := gonfig.GetConf("C:\\cdc-tools\\data sql\\user-level-1.json", &temp)
//	if err != nil {
//		assert.FailNow(t, err.Error())
//	}
//
//	db := connectDB()
//	for i := 0; i < len(temp.Rows); i++ {
//		data := temp.Rows[i]
//		nexchiefAccount, err = dao.GetNexchiefAccountID(db, data.NcCode)
//		if err != nil {
//			continue
//		}
//		data.Level = 1
//		data.Code = data.Code1
//		//err = getParent(db, nexchiefAccount.ID.Int64, &data, nil)
//		//if err != nil {
//		//	continue
//		//}
//		err = getID(db, nexchiefAccount.ID.Int64, &data)
//		if err != nil {
//			continue
//		} else if data.ID > 0 {
//			// update
//			err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		} else {
//			err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		}
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//	}
//}
//
//func TestInsertUserLeve2(t *testing.T) {
//	var (
//		temp            model.DataUserLevel
//		nexchiefAccount model.NexchiefAccount
//	)
//	err := gonfig.GetConf("C:\\cdc-tools\\data sql\\user-level-2.json", &temp)
//	if err != nil {
//		assert.FailNow(t, err.Error())
//	}
//
//	db := connectDB()
//	for i := 0; i < len(temp.Rows); i++ {
//		data := temp.Rows[i]
//		nexchiefAccount, err = dao.GetNexchiefAccountID(db, data.NcCode)
//		if err != nil {
//			continue
//		}
//		data.Level = 2
//		data.Code = data.Code2
//		err = getParent(db, nexchiefAccount.ID.Int64, &data, []interface{}{data.Code1})
//		if err != nil {
//			continue
//		}
//		err = getID(db, nexchiefAccount.ID.Int64, &data)
//		if err != nil {
//			continue
//		} else if data.ID > 0 {
//			// update
//			err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		} else {
//			err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		}
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//	}
//}
//
//func TestInsertUserLeve3(t *testing.T) {
//	var (
//		temp            model.DataUserLevel
//		nexchiefAccount model.NexchiefAccount
//	)
//	err := gonfig.GetConf("C:\\cdc-tools\\data sql\\user-level-3.json", &temp)
//	if err != nil {
//		assert.FailNow(t, err.Error())
//	}
//
//	db := connectDB()
//	for i := 0; i < len(temp.Rows); i++ {
//		data := temp.Rows[i]
//		nexchiefAccount, err = dao.GetNexchiefAccountID(db, data.NcCode)
//		if err != nil {
//			continue
//		}
//		data.Level = 3
//		data.Code = data.Code3
//		err = getParent(db, nexchiefAccount.ID.Int64, &data, []interface{}{data.Code1, data.Code2})
//		if err != nil {
//			continue
//		}
//		err = getID(db, nexchiefAccount.ID.Int64, &data)
//		if err != nil {
//			continue
//		} else if data.ID > 0 {
//			// update
//			err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		} else {
//			err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		}
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//
//	}
//
//}
//
//func TestInsertUserLeve4(t *testing.T) {
//	var (
//		temp            model.DataUserLevel
//		nexchiefAccount model.NexchiefAccount
//	)
//	err := gonfig.GetConf("C:\\cdc-tools\\data sql\\user-level-4.json", &temp)
//	if err != nil {
//		assert.FailNow(t, err.Error())
//	}
//
//	db := connectDB()
//	for i := 0; i < len(temp.Rows); i++ {
//		data := temp.Rows[i]
//		nexchiefAccount, err = dao.GetNexchiefAccountID(db, data.NcCode)
//		if err != nil {
//			continue
//		}
//		data.Level = 4
//		data.Code = data.Code4
//		err = getParent(db, nexchiefAccount.ID.Int64, &data, []interface{}{data.Code1, data.Code2, data.Code3})
//		if err != nil {
//			continue
//		}
//		err = getID(db, nexchiefAccount.ID.Int64, &data)
//		if err != nil {
//			continue
//		} else if data.ID > 0 {
//			// update
//			err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		} else {
//			err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		}
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//
//	}
//
//}
//
//func TestInsertUserLeve5(t *testing.T) {
//	var (
//		temp            model.DataUserLevel
//		nexchiefAccount model.NexchiefAccount
//	)
//	err := gonfig.GetConf("C:\\cdc-tools\\data sql\\user-level-5.json", &temp)
//	if err != nil {
//		assert.FailNow(t, err.Error())
//	}
//
//	db := connectDB()
//	for i := 0; i < len(temp.Rows); i++ {
//		data := temp.Rows[i]
//		nexchiefAccount, err = dao.GetNexchiefAccountID(db, data.NcCode)
//		if err != nil {
//			continue
//		}
//		data.Level = 5
//		data.Code = data.Code5
//		err = getParent(db, nexchiefAccount.ID.Int64, &data, []interface{}{data.Code1, data.Code2, data.Code3, data.Code4})
//		if err != nil {
//			continue
//		}
//		err = getID(db, nexchiefAccount.ID.Int64, &data)
//		if err != nil {
//			continue
//		} else if data.ID > 0 {
//			// update
//			err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		} else {
//			err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
//		}
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//
//	}
//
//}
//func TestQueryParent(t *testing.T) {
//	m := make(map[string]interface{})
//	m["mapping_nexseller_id"] = 1
//	query, param := GetQueryParent("geo_tree", "geo_tree", "code", m, []interface{}{"GEO1", "GEO2", "GEO3", ""}, 0)
//	fmt.Println(query)
//	fmt.Println(len(param))
//	fmt.Println(param)
//}
//
//func TestPointer(t *testing.T) {
//	arr := []string{
//		"satu", "dua", "tiga", "empat",
//	}
//	funcTest(&arr)
//	for i := 0; i < len(arr); i++ {
//		fmt.Println(arr[i])
//	}
//}
//
//func funcTest(arr *[]string) {
//	temp := *arr
//	for i := 0; i < len(temp); i++ {
//		temp[i] = temp[i] + "--"
//	}
//}
