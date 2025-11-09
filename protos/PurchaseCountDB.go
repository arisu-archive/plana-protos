package protos

type PurchaseCountDB struct {
	ShopCashId      int64   `json:",omitempty,omitzero"`
	PurchaseCount   int32   `json:",omitempty,omitzero"`
	ResetDate       MxTime  `json:",omitempty,omitzero"`
	PurchaseDate    *MxTime `json:",omitempty,omitzero"`
	ManualResetDate *MxTime `json:",omitempty,omitzero"`
}
