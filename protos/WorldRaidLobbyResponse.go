package protos

type WorldRaidLobbyResponse struct {
	ResponsePacket
	ClearHistoryDBs     []WorldRaidClearHistoryDB      `json:",omitempty,omitzero"`
	LocalBossDBs        []WorldRaidLocalBossDB         `json:",omitempty,omitzero"`
	BossGroups          map[int64][]WorldRaidBossGroup `json:",omitempty,omitzero"`
	WorldRaidProgressDB WorldRaidProgressDB            `json:",omitempty,omitzero"`
}
