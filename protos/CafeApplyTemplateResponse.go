package protos

type CafeApplyTemplateResponse struct {
	ResponsePacket
	CafeDBs      []CafeDB      `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
