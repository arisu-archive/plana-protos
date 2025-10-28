package protos

type EventContentTreasureFlipRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Round int32 `json:",omitempty,omitzero"`
	Cells []EventContentTreasureCell `json:",omitempty,omitzero"`
}
