package protos

type EliminateRaidSeasonRewardResponse struct {
	ResponsePacket
	ParcelResultDB   ParcelResultDB
	ReceiveRewardIds []int64
}
