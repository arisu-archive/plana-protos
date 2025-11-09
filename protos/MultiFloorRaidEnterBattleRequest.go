package protos

type MultiFloorRaidEnterBattleRequest struct {
	RequestPacket
	SeasonId       int64               `json:",omitempty,omitzero"`
	Difficulty     int32               `json:",omitempty,omitzero"`
	EchelonId      int32               `json:",omitempty,omitzero"`
	AssistUseInfos []ClanAssistUseInfo `json:",omitempty,omitzero"`
}
