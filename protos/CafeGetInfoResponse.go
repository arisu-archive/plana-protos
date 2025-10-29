package protos

type CafeGetInfoResponse struct {
	ResponsePacket
	CafeDB CafeDB `json:",omitempty,omitzero"`
	CafeDBs []CafeDB `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
