package protos

type EventContentSelectBuffRequest struct {
	RequestPacket
	SelectedBuffId int64 `json:",omitempty,omitzero"`
}
