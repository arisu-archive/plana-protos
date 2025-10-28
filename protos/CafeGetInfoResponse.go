package protos

type CafeGetInfoResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDB CafeDB `json:",omitempty,omitzero"`
	CafeDBs []CafeDB `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
