package protos

type MiniGameDefenseGetInfoResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventPointAmount int64 `json:",omitempty,omitzero"`
	DefenseStageHistoryDBs []MiniGameDefenseStageHistoryDB `json:",omitempty,omitzero"`
}
