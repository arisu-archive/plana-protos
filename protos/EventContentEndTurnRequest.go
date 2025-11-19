package protos

type EventContentEndTurnRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	StageUniqueId  int64 `json:",omitempty,omitzero"`
}
