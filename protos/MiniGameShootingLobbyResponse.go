package protos

type MiniGameShootingLobbyResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	HistoryDBs []MiniGameShootingHistoryDB `json:",omitempty,omitzero"`
}
