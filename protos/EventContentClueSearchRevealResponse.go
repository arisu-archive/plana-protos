package protos

type EventContentClueSearchRevealResponse struct {
	ResponsePacket
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
}
