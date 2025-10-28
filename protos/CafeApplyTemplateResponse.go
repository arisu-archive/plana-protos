package protos

type CafeApplyTemplateResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDBs []CafeDB `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
