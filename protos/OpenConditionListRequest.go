package protos

type OpenConditionListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
