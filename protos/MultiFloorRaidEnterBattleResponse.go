package protos

type MultiFloorRaidEnterBattleResponse struct {
	ResponsePacket
	AssistCharacterDBs []AssistCharacterDB `json:",omitempty,omitzero"`
}
