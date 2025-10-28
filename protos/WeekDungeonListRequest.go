package protos

type WeekDungeonListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
