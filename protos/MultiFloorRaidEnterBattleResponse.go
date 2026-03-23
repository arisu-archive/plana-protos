package protos

type MultiFloorRaidEnterBattleResponse struct {
	ResponsePacket
	AssistCharacterDBs []*AssistCharacterDB
}
