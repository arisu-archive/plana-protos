package protos

type WorldRaidLobbyResponse struct {
	ResponsePacket
	ClearHistoryDBs []WorldRaidClearHistoryDB `json:",omitempty,omitzero"`
	LocalBossDBs    []WorldRaidLocalBossDB    `json:",omitempty,omitzero"`
	BossGroups      []WorldRaidBossGroup      `json:",omitempty,omitzero"`
}
