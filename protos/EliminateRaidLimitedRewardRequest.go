package protos

type EliminateRaidLimitedRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
