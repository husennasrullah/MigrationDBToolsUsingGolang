package service
//
//import (
//	"database/sql"
//	"fmt"
//	"github.com/stretchr/testify/assert"
//	"github.com/tkanos/gonfig"
//	"rohmat.co.id/dao"
//	"rohmat.co.id/model"
//	"testing"
//)
//
//func TestInsertSalesman(t *testing.T) {
//	var (
//		temp structSalesman
//	)
//	err := gonfig.GetConf("C:\\cdc-tools\\data sql\\salesman.json", &temp)
//	if err != nil {
//		assert.FailNow(t, err.Error())
//	}
//
//	db := connectDB()
//	for i := 0; i < len(temp.Rows); i++ {
//		data := temp.Rows[i]
//		var (
//			nexchiefAccount       model.NexchiefAccount
//			mnID, personProfileID int64
//		)
//
//		nexchiefAccount, err = dao.GetNexchiefAccountID(db, data.NcCode)
//		if err != nil {
//			continue
//		}
//		// get mapping nexseller
//		mnID, err = dao.GetMappingNexsellerID(db, nexchiefAccount.ID.Int64, data.MnCode)
//		if err != nil {
//			fmt.Println(err.Error())
//			return
//		} else if mnID == 0 {
//			fmt.Println("mapping nexseller not exist :", data.MnCode)
//			continue
//		}
//		// get data FK
//		err = dao.GetSalesmanFKID(db, mnID, &data, nexchiefAccount.Schema.String)
//		if err != nil {
//			fmt.Println("err get Salesman :", err.Error())
//			return
//		} else if data.ID == 0 {
//			// get person profile id
//			personProfileID, err = dao.GetPersonProfileID(db, &data)
//			if err != nil {
//				fmt.Println("err get person profile :", err.Error())
//				return
//			}
//			//else if personProfileID == 0 {
//			//	personProfileID, err = insertPersonProfile(db, &data)
//			//	if err != nil {
//			//		fmt.Println("err insert profile :", err.Error())
//			//		return
//			//	}
//			//}
//			err = dao.InsertSalesman(db, &data, nexchiefAccount, personProfileID, mnID)
//			//err = insertNexsellerProduct(db, nexchiefAccount, mnID, &data)
//			if err != nil {
//				fmt.Println("err insert :", err.Error())
//				return
//			}
//		}
//	}
//
//}
