package protos

type EliminateRaidLimitedRewardResponse struct {
	ResponsePacket
	ParcelResultDB   *ParcelResultDB `json:",omitempty,omitzero"`
	ReceiveRewardIds []int64
}
