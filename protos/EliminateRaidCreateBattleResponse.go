package protos

type EliminateRaidCreateBattleResponse struct {
	ResponsePacket
	RaidDB            RaidDB
	RaidBattleDB      RaidBattleDB
	AccountCurrencyDB AccountCurrencyDB
	AssistCharacterDB AssistCharacterDB
}
