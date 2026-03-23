package protos

type EventListResponse struct {
	ResponsePacket
	EventInfoDBs []*EventInfoDB
}
