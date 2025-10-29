package protos

type BattlePassMissionListResponse struct {
	ResponsePacket
	MissionHistoryUniqueIds []int64 `json:",omitempty,omitzero"`
	ProgressDBs []MissionProgressDB `json:",omitempty,omitzero"`
}
