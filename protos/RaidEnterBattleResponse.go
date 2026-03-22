package protos

type RaidEnterBattleResponse struct {
	ResponsePacket
	RaidDB            RaidDB
	RaidBattleDB      RaidBattleDB
	AccountCurrencyDB AccountCurrencyDB
	AssistCharacterDB AssistCharacterDB
}
