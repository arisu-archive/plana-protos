package protos

type EliminateRaidLimitedRewardResponse struct {
	ResponsePacket
	ParcelResultDB   ParcelResultDB
	ReceiveRewardIds []int64
}
