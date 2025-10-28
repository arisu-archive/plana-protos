package protos

type EliminateRaidLobbyInfoDB struct {
	RaidLobbyInfoDB
	OpenedBossGroups []string `json:",omitempty,omitzero"`
	BestRankingPointPerBossGroup map[string]int64 `json:",omitempty,omitzero"`
}
