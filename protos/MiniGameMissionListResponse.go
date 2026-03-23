package protos

type MiniGameMissionListResponse struct {
	ResponsePacket
	MissionHistoryUniqueIds  []int64
	ProgressDBs              []*MissionProgressDB
	ClearedOrignalMissionIds []int64
}
