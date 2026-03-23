package protos

type ClearDeckListResponse struct {
	ResponsePacket
	ClearDeckDBs []*ClearDeckDB
}
