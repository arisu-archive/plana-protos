package protos

type MiniGameStageListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MiniGameHistoryDBs []MiniGameHistoryDB `json:",omitempty,omitzero"`
}
