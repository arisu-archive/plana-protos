package protos

type BillingPurchaseListByYostarResponse struct {
	ResponsePacket
	CountList []PurchaseCountDB `json:",omitempty,omitzero"`
	OrderList []PurchaseOrderDB `json:",omitempty,omitzero"`
	MonthlyProductList []MonthlyProductPurchaseDB `json:",omitempty,omitzero"`
	BlockedProductDBs []BlockedProductDB `json:",omitempty,omitzero"`
	BattlePassProductList []BattlePassProductPurchaseDB `json:",omitempty,omitzero"`
}
