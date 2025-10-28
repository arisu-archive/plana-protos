package protos

type EliminateRaidLimitedRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ReceiveRewardIds []int64 `json:",omitempty,omitzero"`
}
