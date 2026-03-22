package protos

type ReceiveAccountLevelRewardResponse struct {
	ResponsePacket
	ReceivedAccountLevelRewardIds []int64
	ParcelResultDB                ParcelResultDB
}
