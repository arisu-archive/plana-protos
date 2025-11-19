package protos

type OpenConditionEventListResponse struct {
	ResponsePacket
	ConquestTiles         map[int64][]ConquestTileDB       `json:",omitempty,omitzero"`
	WorldRaidLocalBossDBs map[int64][]WorldRaidLocalBossDB `json:",omitempty,omitzero"`
}
