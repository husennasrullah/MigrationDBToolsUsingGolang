package model

type PurchaseReceiveItem struct {
	ID                 int64  `json:"id"`
	NcCode             string `json:"principalID"`
	MnCode             string `json:"distributorID"`
	PurchaseReceiveID  int64
	ProductID          int64
	DocumentNumber     string  `json:"documentNumber"`
	ProductCode        string  `json:"productCode"`
	BatchNumber        string  `json:"batchNumber"`
	QtyReceived        int     `json:"qtyReceived"`
	QtyFreeGood        int     `json:"qtyFreeGood"`
	BuyingPrice        float64 `json:"buyingPrice"`
	Conversion         int     `json:"conversion"`
	LineDiscount1      float64 `json:"lineDiscount1"`
	LineDiscount2      float64 `json:"lineDiscount2"`
	LineDiscount3      float64 `json:"lineDiscount3"`
	LineDiscount4      float64 `json:"lineDiscount4"`
	LineDiscount5      float64 `json:"lineDiscount5"`
	LineDiscountBased  float64 `json:"lineDiscountBased"`
	LineDiscountAmount float64 `json:"lineDiscountAmount"`
	DiscountAmount     float64 `json:"discountAmount"`
	OtherCostBeforeTax float64 `json:"otherCostBeforeTax"`
	TaxAmount          float64 `json:"taxAmount"`
	LineGrossAmount    float64 `json:"lineGrossAmount"`
	LineNetAmount      float64 `json:"lineNetAmount"`
}
