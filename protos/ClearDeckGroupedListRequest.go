package protos

type ClearDeckGroupedListRequest struct {
	RequestPacket
	ClearDeckKey *ClearDeckKey `json:",omitempty,omitzero"`
}
