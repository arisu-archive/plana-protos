package protos

type ClearDeckListRequest struct {
	RequestPacket
	ClearDeckKey ClearDeckKey `json:",omitempty,omitzero"`
}
