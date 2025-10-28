package protos

type CraftShiftingCompleteProcessAllRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
