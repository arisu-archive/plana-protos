package protos

type BillingTransactionEndByYostarResponse struct {
	ResponsePacket
	ParcelResult          ParcelResultDB
	MailDB                MailDB
	CountList             []PurchaseCountDB
	PurchaseCount         int32 `json:",omitempty,omitzero"`
	MonthlyProductList    []MonthlyProductPurchaseDB
	BattlePassInfo        BattlePassInfoDB
	BattlePassProductList []BattlePassProductPurchaseDB
}
