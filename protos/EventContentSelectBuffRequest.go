package protos

type EventContentSelectBuffRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SelectedBuffId int64 `json:",omitempty,omitzero"`
}
