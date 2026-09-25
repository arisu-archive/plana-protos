package protos

type EventContentClueSearchBatchSubmitResponse struct {
	ResponsePacket
	SaveDB         *ClueSearchSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB *ParcelResultDB   `json:",omitempty,omitzero"`
}
