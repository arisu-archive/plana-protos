package protos

type CafeRenamePresetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
