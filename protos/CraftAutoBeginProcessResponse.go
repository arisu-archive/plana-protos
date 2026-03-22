package protos

type CraftAutoBeginProcessResponse struct {
	ResponsePacket
	CraftInfoDBs   []CraftInfoDB
	ParcelResultDB ParcelResultDB
}
