package protos

type BattlePassMissionSingleRewardResponse struct {
	ResponsePacket
	AddedHistoryDB MissionHistoryDB
	ParcelResultDB ParcelResultDB
	BattlePassInfo BattlePassInfoDB
}
