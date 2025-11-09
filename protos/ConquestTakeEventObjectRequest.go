package protos

type ConquestTakeEventObjectRequest struct {
	RequestPacket
	EventContentId     int64 `json:",omitempty,omitzero"`
	ConquestObjectDBId int64 `json:",omitempty,omitzero"`
}
