package protos

type EventRewardIncreaseRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
