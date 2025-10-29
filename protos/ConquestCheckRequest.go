package protos

type ConquestCheckRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
