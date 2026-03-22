package protos

type MissionListResponse struct {
	ResponsePacket
	MissionHistoryUniqueIds  []int64
	ProgressDBs              []MissionProgressDB
	DailySuddenMissionInfo   MissionInfo
	ClearedOrignalMissionIds []int64
}
