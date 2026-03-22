package protos

type EventContentPermanentListResponse struct {
	ResponsePacket
	PermanentDBs []EventContentPermanentDB
}
