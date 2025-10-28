package protos

type TimeAttackDungeonEnterBattleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RoomId int64 `json:",omitempty,omitzero"`
	AssistUseInfo ClanAssistUseInfo `json:",omitempty,omitzero"`
}
