package protos

type AccountLinkRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
