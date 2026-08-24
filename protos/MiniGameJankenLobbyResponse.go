package protos

type MiniGameJankenLobbyResponse struct {
	ResponsePacket
	SaveDB          *MiniGameJankenSaveDB `json:",omitempty,omitzero"`
	StageHistoryDBs []*MiniGameJankenStageHistoryDB
}
