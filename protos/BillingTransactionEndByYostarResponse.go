package protos

type BillingTransactionEndByYostarResponse struct {
	ResponsePacket
	ParcelResult          *ParcelResultDB `json:",omitempty,omitzero"`
	MailDB                *MailDB         `json:",omitempty,omitzero"`
	CountList             []*PurchaseCountDB
	PurchaseCount         int32 `json:",omitempty,omitzero"`
	MonthlyProductList    []*MonthlyProductPurchaseDB
	BattlePassInfo        *BattlePassInfoDB `json:",omitempty,omitzero"`
	BattlePassProductList []*BattlePassProductPurchaseDB
}
