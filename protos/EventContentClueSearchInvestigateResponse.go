package protos

type EventContentClueSearchInvestigateResponse struct {
	ResponsePacket
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
}
