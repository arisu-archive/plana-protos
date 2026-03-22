package protos

type CafeApplyPresetResponse struct {
	ResponsePacket
	CafeDBs      []CafeDB
	FurnitureDBs []FurnitureDB
}
