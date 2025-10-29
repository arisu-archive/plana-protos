package protos

type ConquestMainStoryCheckRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
