package protos

type PermanentRaidLobbyResponse struct {
	ResponsePacket
	BossManageDBs       []PermanentRaidBossManageDB
	BestScoreHistoryDBs []PermanentRaidBestScoreHistoryDB
	BattleHistoryDB     *PermanentRaidBattleHistoryDB `json:",omitempty,omitzero"`
}
