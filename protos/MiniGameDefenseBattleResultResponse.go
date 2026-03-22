package protos

type MiniGameDefenseBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	StageHistoryDB MiniGameDefenseStageHistoryDB
}
