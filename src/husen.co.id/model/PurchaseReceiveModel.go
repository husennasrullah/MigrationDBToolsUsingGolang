package model

import "time"

type PurchaseReceive struct {
	ID                     int64  `json:"id"`
	NcCode                 string `json:"principalID"`
	MnCode                 string `json:"distributorID"`
	DocumentNumber         string `json:"documentNumber"`
	TransactionDateStr     string `json:"transactionDate"`
	TransactionDate        time.Time
	DeliveryDocumentNumber string `json:"deliveryDocumentNumber"`
	DeliveryDateStr        string `json:"deliveryDate"`
	DeliveryDate           time.Time
	ManualPONumber         string `json:"manualPONumber"`
	TotalGrossAmount       int    `json:"totalGrossAmount"`
	TotalNetAmount         int    `json:"totalNetAmount"`
	InvoiceNumber          string `json:"invoiceNumber"`
	InvoiceDateStr         string `json:"invoiceDate"`
	InvoiceDate            time.Time
	LastSync               interface{} `json:"lastSync"`
	Version                interface{} `json:"version"`
	TotalLineDiscount      int         `json:"totalLineDiscount"`
	TotalDiscount          int         `json:"totalDiscount"`
	TotalCostBeforeTax     int         `json:"totalCostBeforeTax"`
	TotalCostAfterTax      int         `json:"totalCostAfterTax"`
	TotalTax               int         `json:"totalTax"`
	WarehouseID            string      `json:"warehouseID"`
	WarehouseIDDB          int64       `json:"warehouse_iddb"`
	Created                interface{} `json:"created"`
	Modified               interface{} `json:"modified"`
	ModifiedBy             string      `json:"modifiedBy"`
	Status                 string      `json:"status"`
	TransitDocumentNumber  string      `json:"transitDocumentNumber"`
	Statuspowerbi          string      `json:"statuspowerbi"`
	Sentpowerbi            interface{} `json:"sentpowerbi"`
}
