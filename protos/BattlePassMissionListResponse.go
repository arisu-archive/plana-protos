package protos

type BattlePassMissionListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MissionHistoryUniqueIds []int64 `json:",omitempty,omitzero"`
	ProgressDBs []MissionProgressDB `json:",omitempty,omitzero"`
}
