package protos

type PermanentRaidLobbyResponse struct {
	ResponsePacket
	BossManageDBs       []PermanentRaidBossManageDB       `json:",omitempty,omitzero"`
	BestScoreHistoryDBs []PermanentRaidBestScoreHistoryDB `json:",omitempty,omitzero"`
	BattleHistoryDB     PermanentRaidBattleHistoryDB      `json:",omitempty,omitzero"`
}
