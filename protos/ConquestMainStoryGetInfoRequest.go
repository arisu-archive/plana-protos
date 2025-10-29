package protos

type ConquestMainStoryGetInfoRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
