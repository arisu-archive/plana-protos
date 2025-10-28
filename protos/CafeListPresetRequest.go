package protos

type CafeListPresetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
