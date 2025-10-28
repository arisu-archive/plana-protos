package protos

type EventContentConcentrationFlipCardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	FirstIndex int32 `json:",omitempty,omitzero"`
	SecondIndex int32 `json:",omitempty,omitzero"`
}
