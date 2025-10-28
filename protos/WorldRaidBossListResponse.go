package protos

type WorldRaidBossListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BossListInfoDBs []WorldRaidBossListInfoDB `json:",omitempty,omitzero"`
}
