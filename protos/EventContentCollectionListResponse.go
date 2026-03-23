package protos

type EventContentCollectionListResponse struct {
	ResponsePacket
	EventContentUnlockCGDBs []*EventContentCollectionDB
}
