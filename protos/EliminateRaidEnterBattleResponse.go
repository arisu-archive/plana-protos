package protos

type EliminateRaidEnterBattleResponse struct {
	ResponsePacket
	RaidDB            RaidDB
	RaidBattleDB      RaidBattleDB
	AccountCurrencyDB AccountCurrencyDB
	AssistCharacterDB AssistCharacterDB
}
