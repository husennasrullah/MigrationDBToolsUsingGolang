package service

import (
	"database/sql"
	"encoding/json"
	"husen.co.id/dao"
	"husen.co.id/model"
	"testing"
	"time"
)

func TestSavePurchaseReceiveItem(t *testing.T) {
	// 3 eror
	path := "D:\\GO APP\\MigrationDB\\purchase_receive_item.json"
	StartReadFile(path, SavePurchaseReceiveItem)
	time.Sleep(5 * time.Second)
}

func SavePurchaseReceiveItem(db *sql.DB, dataByte []byte) (errorModel model.ErrorModel) {
	var (
		nexchiefAccount model.NexchiefAccount
		data            model.PurchaseReceiveItem
		mnID int64
	)

	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, errorModel = dao.GetNexchiefAccountID(db, data.NcCode)
	if errorModel.Error != nil {
		return
	}

	// get mapping nexseller
	mnID, errorModel = dao.GetMappingNexsellerID(db, nexchiefAccount.ID.Int64, data.MnCode)
	if errorModel.Error != nil {
		return
	}

	// get data FK
	errorModel = dao.GetPurchaseReceiveItemFK(db, nexchiefAccount, &data, mnID)
	if errorModel.Error != nil {
		return
	} else if data.ID == 0 {
		errorModel = dao.InsertPurchaseReceiveItem(db, nexchiefAccount, mnID, &data)
		if errorModel.Error != nil {
			return
		}
	}
	return
}
