package protos

type BillingTransactionEndByYostarResponse struct {
	ResponsePacket
	ParcelResult          ParcelResultDB                `json:",omitempty,omitzero"`
	MailDB                MailDB                        `json:",omitempty,omitzero"`
	CountList             []PurchaseCountDB             `json:",omitempty,omitzero"`
	PurchaseCount         int32                         `json:",omitempty,omitzero"`
	MonthlyProductList    []MonthlyProductPurchaseDB    `json:",omitempty,omitzero"`
	BattlePassInfo        BattlePassInfoDB              `json:",omitempty,omitzero"`
	BattlePassProductList []BattlePassProductPurchaseDB `json:",omitempty,omitzero"`
}
