package protos

type CafeApplyTemplateResponse struct {
	ResponsePacket
	CafeDBs      []*CafeDB
	FurnitureDBs []*FurnitureDB
}
