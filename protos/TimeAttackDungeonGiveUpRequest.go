package protos

type TimeAttackDungeonGiveUpRequest struct {
	RequestPacket
	RoomId int64 `json:",omitempty,omitzero"`
}
