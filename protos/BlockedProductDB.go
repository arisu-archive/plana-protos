package protos

type BlockedProductDB struct {
	CashProductId   int64             `json:",omitempty,omitzero"`
	MarketBlockType ShopCashBlockType `json:",omitempty,omitzero"`
	BeginDate       MxTime            `json:",omitempty,omitzero"`
	EndDate         MxTime            `json:",omitempty,omitzero"`
}
