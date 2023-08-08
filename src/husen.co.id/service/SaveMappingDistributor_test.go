package service
//
//import (
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	"github.com/stretchr/testify/assert"
//	"github.com/tkanos/gonfig"
//	"rohmat.co.id/dao"
//	"rohmat.co.id/model"
//	"testing"
//)
//
//func TestInsertMappingNexseller(t *testing.T) {
//	var (
//		temp            structMappingDistributor
//		nexchiefAccount model.NexchiefAccount
//	)
//	err := gonfig.GetConf("C:\\cdc-tools\\data sql\\mapping-distributor.json", &temp)
//	if err != nil {
//		assert.FailNow(t, err.Error())
//	}
//
//	db := connectDB()
//
//	for i := 0; i < len(temp.Rows); i++ {
//		data := temp.Rows[i]
//		nexchiefAccount, err = dao.GetNexchiefAccountID(db, data.NcCode)
//		if err != nil {
//			continue
//		}
//		//err = getParent(db, nexchiefAccount.ID.Int64, &data, nil)
//		//if err != nil {
//		//	continue
//		//}
//		err = dao.GetMappingNexseller(db, nexchiefAccount.ID.Int64, &data)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		} else if data.ID == 0 {
//			// insert
//			err = dao.InsertMappingNexseller(db, nexchiefAccount.ID.Int64, &data)
//		}
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//	}
//}
