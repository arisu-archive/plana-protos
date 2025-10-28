package protos

type RaidBattleUpdateRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidServerId int64 `json:",omitempty,omitzero"`
	RaidBossIndex int32 `json:",omitempty,omitzero"`
	CumulativeDamage int64 `json:",omitempty,omitzero"`
	CumulativeGroggyPoint int64 `json:",omitempty,omitzero"`
}
