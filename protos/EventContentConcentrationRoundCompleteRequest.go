package protos

type EventContentConcentrationRoundCompleteRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
