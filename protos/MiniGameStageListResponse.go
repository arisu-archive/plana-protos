package protos

type MiniGameStageListResponse struct {
	ResponsePacket
	MiniGameHistoryDBs []*MiniGameHistoryDB
}
