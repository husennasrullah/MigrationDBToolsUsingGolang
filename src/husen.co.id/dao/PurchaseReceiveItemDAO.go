package dao

import (
	"database/sql"
	"husen.co.id/model"
)

func GetPurchaseReceiveItemFK(db *sql.DB, nc model.NexchiefAccount, data *model.PurchaseReceiveItem, mnID int64) (errModel model.ErrorModel) {
	query :=
		"SELECT 'pr', id FROM " + getSchema("nexseller_purchase_receive", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND document_number = $2 " +
			" UNION ALL " +
			"SELECT 'np', id FROM " + getSchema("nexseller_product", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND product_code= $3 "

	param := []interface{}{
		mnID, data.DocumentNumber, data.ProductCode,
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
			case "pri":
				data.ID = id.Int64
			case "pr":
				data.PurchaseReceiveID = id.Int64
			case "np":
				data.ProductID = id.Int64
			}
		}
	}
	return
}

func InsertPurchaseReceiveItem(db *sql.DB, nc model.NexchiefAccount, mnID int64, data *model.PurchaseReceiveItem) (err model.ErrorModel) {
	query :=
		"INSERT INTO " +
			getSchema("nexseller_purchase_receive_item", nc.Schema.String) +
			"(nexchief_account_id, mapping_nexseller_id, nexseller_purchase_receive_id, " +
			"nexseller_product_id, batch_number, qty_received, " +
			"qty_free_good, buying_price, conversion, " +
			"line_discount_1, line_discount_2, line_discount_3, " +
			"line_discount_4, line_discount_5, line_discount_based, " +
			"line_discount_amount, discount_amount, other_cost_before_tax, " +
			"tax_amount, line_gross_amount, line_net_amount) " +
			"VALUES" +
			"($1, $2, $3, " +
			"$4, $5, $6, " +
			"$7, $8, $9, " +
			"$10, $11, $12, " +
			"$13, $14, $15, " +
			"$16, $17, $18, " +
			"$19, $20, $21) " +
			"RETURNING id;"

	param := []interface{}{
		nc.ID.Int64, mnID, data.PurchaseReceiveID,
		data.ProductID, data.BatchNumber, data.QtyReceived,
		data.QtyFreeGood, data.BuyingPrice, data.Conversion,
		data.LineDiscount1, data.LineDiscount2, data.LineDiscount3,
		data.LineDiscount4, data.LineDiscount5, data.LineDiscountBased,
		data.LineDiscountAmount, data.DiscountAmount, data.OtherCostBeforeTax,
		data.TaxAmount, data.LineGrossAmount, data.LineNetAmount,
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
