package protos

type TimeAttackDungeonGiveUpRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RoomId int64 `json:",omitempty,omitzero"`
}
