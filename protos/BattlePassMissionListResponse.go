package protos

type BattlePassMissionListResponse struct {
	ResponsePacket
	MissionHistoryUniqueIds []int64
	ProgressDBs             []*MissionProgressDB
}
