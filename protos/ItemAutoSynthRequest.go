package protos

type ItemAutoSynthRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetParcels []ParcelKeyPair `json:",omitempty,omitzero"`
}
