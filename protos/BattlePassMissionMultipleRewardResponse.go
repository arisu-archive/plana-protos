package protos

type BattlePassMissionMultipleRewardResponse struct {
	ResponsePacket
	AddedHistoryDBs []MissionHistoryDB
	ParcelResultDB  ParcelResultDB
	BattlePassInfo  BattlePassInfoDB
}
