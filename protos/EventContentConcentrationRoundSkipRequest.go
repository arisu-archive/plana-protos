package protos

type EventContentConcentrationRoundSkipRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
