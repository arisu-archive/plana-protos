package protos

type ClearDeckListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClearDeckKey ClearDeckKey `json:",omitempty,omitzero"`
}
