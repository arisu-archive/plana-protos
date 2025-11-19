package protos

type MissionListResponse struct {
	ResponsePacket
	MissionHistoryUniqueIds  []int64             `json:",omitempty,omitzero"`
	ProgressDBs              []MissionProgressDB `json:",omitempty,omitzero"`
	DailySuddenMissionInfo   MissionInfo         `json:",omitempty,omitzero"`
	ClearedOrignalMissionIds []int64             `json:",omitempty,omitzero"`
}
