package protos

type EventContentCollectionListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentUnlockCGDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
