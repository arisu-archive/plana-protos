package protos

type MiniGameShootingLobbyResponse struct {
	ResponsePacket
	HistoryDBs []MiniGameShootingHistoryDB `json:",omitempty,omitzero"`
}
