package protos

type PermanentRaidEnterBattleResponse struct {
	ResponsePacket
	BattleHistoryDB   PermanentRaidBattleHistoryDB
	AssistCharacterDB AssistCharacterDB
}
