package protos

type MissionListResponse struct {
	ResponsePacket
	MissionHistoryUniqueIds  []int64
	ProgressDBs              []MissionProgressDB
	DailySuddenMissionInfo   *MissionInfo `json:",omitempty,omitzero"`
	ClearedOrignalMissionIds []int64
}
