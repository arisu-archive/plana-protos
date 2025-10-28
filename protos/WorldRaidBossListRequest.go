package protos

type WorldRaidBossListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SeasonId int64 `json:",omitempty,omitzero"`
	RequestOnlyWorldBossData bool `json:",omitempty,omitzero"`
}
