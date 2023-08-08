package service

import (
	"database/sql"
	"encoding/json"
	"husen.co.id/dao"
	"husen.co.id/model"
	"testing"
	"time"
)

func TestSavePurchaseReceive(t *testing.T) {
	// 3 eror
	path := "D:\\GO APP\\MigrationDB\\purchase_receive.json"
	StartReadFile(path, SavePurchaseReceive)
	time.Sleep(5 * time.Second)
}

func SavePurchaseReceive(db *sql.DB, dataByte []byte) (errorModel model.ErrorModel) {
	var (
		nexchiefAccount model.NexchiefAccount
		data            model.PurchaseReceive
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

	ChangeData(&data)

	// get data FK
	errorModel = dao.GetPurchaseReceiveFK(db, nexchiefAccount, &data, mnID)
	if errorModel.Error != nil {
		return
	} else if data.ID == 0 {
		errorModel = dao.InsertPurchaseReceive(db, nexchiefAccount, mnID, &data)
		if errorModel.Error != nil {
			return
		}
	}
	return
}

func ChangeData (data *model.PurchaseReceive) {
	data.TransactionDate, _ = time.Parse("2006-01-02T15:04:05", data.TransactionDateStr)
	data.InvoiceDate, _ = time.Parse("2006-01-02T15:04:05", data.InvoiceDateStr)
	data.DeliveryDate, _ = time.Parse("2006-01-02T15:04:05", data.DeliveryDateStr)

}
