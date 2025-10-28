package protos

type RaidCompleteListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
