package protos

type CraftCompleteProcessAllRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
