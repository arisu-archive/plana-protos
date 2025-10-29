package protos

type WeekDungeonRetreatRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
