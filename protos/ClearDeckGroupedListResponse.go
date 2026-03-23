package protos

type ClearDeckGroupedListResponse struct {
	ResponsePacket
	ClearDeckGroupedDBs [][]*ClearDeckDB
}
