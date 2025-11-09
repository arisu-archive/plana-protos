package protos

type BattlePassProductPurchaseDB struct {
	ProductId                 int64 `json:",omitempty,omitzero"`
	BattlePassId              int64 `json:",omitempty,omitzero"`
	PurchaseBattlePassGroupId int64 `json:",omitempty,omitzero"`
}
