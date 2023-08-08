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
//
//func TestInsertCompanyProfile(t *testing.T) {
//	path := "C:\\cdc-tools\\data sql\\distributor.json"
//
//	db := connectDB()
//	for i := 0; i < len(temp.Rows); i++ {
//		data := temp.Rows[i]
//		// get company profile
//		var (
//			cpID, parentID int64
//			resultDB       []model.MappingDistributor
//			err error
//		)
//		cpID, err = dao.GetIDCompanyProfile(db, &data)
//		if err != nil {
//			fmt.Println("err company profile : ", err.Error())
//			continue
//		} else if cpID == 0 {
//			// insert
//			cpID, err = dao.InsertCompanyProfile(db, &data)
//			if err != nil {
//				fmt.Println("insert company profile : ", err.Error())
//				continue
//			}
//		}
//		resultDB, err = dao.GetDistributor(db, &data)
//		if err != nil {
//			fmt.Println("get distributor : ", err.Error())
//			continue
//		}
//
//		for _, m := range resultDB {
//			parentID, err = dao.GetParentID(db, &m)
//			if err != nil {
//				fmt.Println("get parent id : ", err.Error())
//				continue
//			}
//			err = dao.UpdateMappingNexseller(db, m.ID, cpID, parentID)
//			if err != nil {
//				fmt.Println("update mapping nexseller : ", err.Error())
//				continue
//			}
//		}
//	}
//}
