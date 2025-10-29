package protos

type EventContentPermanentListResponse struct {
	ResponsePacket
	PermanentDBs []EventContentPermanentDB `json:",omitempty,omitzero"`
}
