package protos

type RaidSeasonRewardResponse struct {
	ResponsePacket
	ParcelResultDB   ParcelResultDB
	ReceiveRewardIds []int64
}
