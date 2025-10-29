package protos

type TimeAttackDungeonEndBattleRequest struct {
	RequestPacket
	EchelonId int32 `json:",omitempty,omitzero"`
	RoomId int64 `json:",omitempty,omitzero"`
	Summary BattleSummary `json:",omitempty,omitzero"`
	AssistUseInfo ClanAssistUseInfo `json:",omitempty,omitzero"`
}
