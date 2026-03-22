package protos

type WorldRaidEnterBattleResponse struct {
	ResponsePacket
	RaidBattleDB       RaidBattleDB
	AssistCharacterDBs []AssistCharacterDB
}
