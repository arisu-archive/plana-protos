package protos

type ReceiveAccountLevelRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
