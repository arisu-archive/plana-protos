package protos

type MiniGameMissionRewardResponse struct {
	ResponsePacket
	AddedHistoryDB MissionHistoryDB
	ParcelResultDB ParcelResultDB
}
