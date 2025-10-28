package protos

type CharacterWeaponExpGrowthRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetCharacterServerId int64 `json:",omitempty,omitzero"`
	ConsumeUniqueIdAndCounts map[int64]int64 `json:",omitempty,omitzero"`
}
