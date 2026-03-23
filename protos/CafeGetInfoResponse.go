package protos

type CafeGetInfoResponse struct {
	ResponsePacket
	CafeDB       *CafeDB `json:",omitempty,omitzero"`
	CafeDBs      []CafeDB
	FurnitureDBs []FurnitureDB
}
