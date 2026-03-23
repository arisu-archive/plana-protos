package protos

type ResetableContentGetResponse struct {
	ResponsePacket
	ResetableContentValueDBs []*ResetableContentValueDB
}
