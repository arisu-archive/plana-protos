package protos

type CafeOpenResponse struct {
	ResponsePacket
	OpenedCafeDB CafeDB        `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
