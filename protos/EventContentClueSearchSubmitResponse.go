package protos

type EventContentClueSearchSubmitResponse struct {
	ResponsePacket
	SaveDB         ClueSearchSaveDB
	ParcelResultDB ParcelResultDB
}
