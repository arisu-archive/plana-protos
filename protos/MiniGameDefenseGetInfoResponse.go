package protos

type MiniGameDefenseGetInfoResponse struct {
	ResponsePacket
	EventPointAmount       int64                           `json:",omitempty,omitzero"`
	DefenseStageHistoryDBs []MiniGameDefenseStageHistoryDB `json:",omitempty,omitzero"`
}
