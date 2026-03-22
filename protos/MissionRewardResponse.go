package protos

type MissionRewardResponse struct {
	ResponsePacket
	AddedHistoryDB MissionHistoryDB
	ParcelResultDB ParcelResultDB
}
