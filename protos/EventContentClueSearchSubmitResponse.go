package protos

type EventContentClueSearchSubmitResponse struct {
	ResponsePacket
	SaveDB         *ClueSearchSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB *ParcelResultDB   `json:",omitempty,omitzero"`
}
