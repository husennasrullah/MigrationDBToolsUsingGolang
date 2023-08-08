package dao

import (
	"database/sql"
	"husen.co.id/model"
)

func GetPurchaseReceiveFK(db *sql.DB, nc model.NexchiefAccount, data *model.PurchaseReceive, mnID int64) (errModel model.ErrorModel) {
	query := "SELECT 'pr', id FROM " + getSchema("nexseller_purchase_receive", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND document_number = $2 " +
		" UNION ALL " +
		"SELECT 'nw', id FROM " + getSchema("nexseller_warehouse", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code= $3 "

	param := []interface{}{
		mnID, data.DocumentNumber, data.WarehouseID,
	}

	rows, err := db.Query(query, param...)
	if err != nil {
		return generateErrorModel(err)
	}
	if rows != nil {
		defer func() {
			err = rows.Close()
			if err != nil {
				errModel = generateErrorModel(err)
				return
			}
		}()

		for rows.Next() {
			var key string
			var id sql.NullInt64

			err = rows.Scan(&key, &id)
			if err != nil {
				return generateErrorModel(err)
			}
			switch key {
			case "pr":
				data.ID = id.Int64
			case "nw":
				data.WarehouseIDDB = id.Int64
			}
		}
	}
	return
}

func InsertPurchaseReceive(db *sql.DB, nc model.NexchiefAccount, mnID int64, data *model.PurchaseReceive) (err model.ErrorModel) {
	query :=
		"INSERT INTO " +
			getSchema("nexseller_purchase_receive", nc.Schema.String) + " " +
			"(nexchief_account_id, mapping_nexseller_id, " +
			"transaction_date, document_number, delivery_document_number, " +
			"delivery_date, manual_po_number, total_gross_amount, " +
			"total_net_amount, invoice_number, invoice_date, " +
			"total_line_discount, total_discount, total_cost_before_tax, " +
			"total_cost_after_tax, total_tax, nexseller_warehouse_id, " +
			"status, transit_document_number) " +
			"VALUES " +
			"($1, $2, $3, " +
			"$4, $5, $6, " +
			"$7, $8, $9, " +
			"$10, $11, $12, " +
			"$13, $14, $15, " +
			"$16, $17, $18, $19) " +
			"RETURNING id;"

	param := []interface{}{
		nc.ID, mnID,
		data.TransactionDate, data.DocumentNumber, data.DeliveryDocumentNumber,
		data.DeliveryDate, data.ManualPONumber, data.TotalGrossAmount,
		data.TotalNetAmount, data.InvoiceNumber, data.InvoiceDate,
		data.TotalLineDiscount, data.TotalDiscount, data.TotalCostBeforeTax,
		data.TotalCostAfterTax, data.TotalTax, data.WarehouseIDDB,
		data.Status, data.TransitDocumentNumber,
	}
	stmt, errS := db.Prepare(query)
	if errS != nil {
		err = generateErrorModel(errS)
		return err
	}
	defer stmt.Close()
	_, errS = stmt.Exec(param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
