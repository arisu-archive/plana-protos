package protos

type EventContentClueSearchRoundCompleteResponse struct {
	ResponsePacket
	SaveDB         ClueSearchSaveDB
	ParcelResultDB ParcelResultDB
}
