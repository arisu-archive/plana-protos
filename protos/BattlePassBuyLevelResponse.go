package protos

type BattlePassBuyLevelResponse struct {
	ResponsePacket
	BattlePassInfo BattlePassInfoDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
}
