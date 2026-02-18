package protos

type EventContentClueSearchRoundCompleteResponse struct {
	ResponsePacket
	SaveDB         ClueSearchSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB   `json:",omitempty,omitzero"`
}
