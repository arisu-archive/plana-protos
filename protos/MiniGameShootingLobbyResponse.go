package protos

type MiniGameShootingLobbyResponse struct {
	ResponsePacket
	HistoryDBs []*MiniGameShootingHistoryDB
}
