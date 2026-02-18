package protos

type EventContentClueSearchGetInfoResponse struct {
	ResponsePacket
	SaveDB ClueSearchSaveDB `json:",omitempty,omitzero"`
}
