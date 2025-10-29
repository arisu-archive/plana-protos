package protos

type EventContentFortuneGachaPurchaseRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
