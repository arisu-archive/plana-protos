package protos

type EventContentTreasureNextRoundRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	Round int32 `json:",omitempty,omitzero"`
}
