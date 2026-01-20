package protos

type AccountLimitedFlashSaleDB struct {
	ShopExcelId int64  `json:",omitempty,omitzero"`
	SaleFrom    MxTime `json:",omitempty,omitzero"`
	SaleTo      MxTime `json:",omitempty,omitzero"`
}
