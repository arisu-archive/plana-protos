package protos

type MiniGameStageListResponse struct {
	ResponsePacket
	MiniGameHistoryDBs []MiniGameHistoryDB `json:",omitempty,omitzero"`
}
