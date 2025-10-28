package protos

type RaidEndBattleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EchelonId int32 `json:",omitempty,omitzero"`
	RaidServerId int64 `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
	Summary BattleSummary `json:",omitempty,omitzero"`
	AssistUseInfo ClanAssistUseInfo `json:",omitempty,omitzero"`
}
