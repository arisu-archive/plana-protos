package protos

type BattlePassBuyLevelResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BattlePassInfo BattlePassInfoDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
}
