package protos

type MiniGameMissionListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MissionHistoryUniqueIds []int64 `json:",omitempty,omitzero"`
	ProgressDBs []MissionProgressDB `json:",omitempty,omitzero"`
	ClearedOrignalMissionIds []int64 `json:",omitempty,omitzero"`
}
