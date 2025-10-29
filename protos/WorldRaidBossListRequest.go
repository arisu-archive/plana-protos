package protos

type WorldRaidBossListRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
	RequestOnlyWorldBossData bool `json:",omitempty,omitzero"`
}
