package protos

type BattlePassBuyLevelResponse struct {
	ResponsePacket
	BattlePassInfo    BattlePassInfoDB
	AccountCurrencyDB AccountCurrencyDB
}
