package protos

type MultiFloorRaidEnterBattleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AssistCharacterDBs []AssistCharacterDB `json:",omitempty,omitzero"`
}
