package protos

type TimeAttackDungeonEnterBattleRequest struct {
	RequestPacket
	RoomId int64 `json:",omitempty,omitzero"`
	AssistUseInfo ClanAssistUseInfo `json:",omitempty,omitzero"`
}
