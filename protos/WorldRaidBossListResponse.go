package protos

type WorldRaidBossListResponse struct {
	ResponsePacket
	BossListInfoDBs []WorldRaidBossListInfoDB `json:",omitempty,omitzero"`
}
