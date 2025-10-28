package protos

type CafeTrophyHistoryRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
