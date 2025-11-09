package protos

type EliminateRaidLimitedRewardResponse struct {
	ResponsePacket
	ParcelResultDB   ParcelResultDB `json:",omitempty,omitzero"`
	ReceiveRewardIds []int64        `json:",omitempty,omitzero"`
}
