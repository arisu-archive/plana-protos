package protos

type CafeApplyPresetResponse struct {
	ResponsePacket
	CafeDBs      []CafeDB      `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
