package protos

type CafeOpenResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	OpenedCafeDB CafeDB `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
