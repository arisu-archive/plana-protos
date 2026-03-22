package protos

type MissionMultipleRewardResponse struct {
	ResponsePacket
	AddedHistoryDBs []MissionHistoryDB
	ParcelResultDB  ParcelResultDB
}
