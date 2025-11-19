package protos

type CharacterWeaponExpGrowthRequest struct {
	RequestPacket
	TargetCharacterServerId  int64           `json:",omitempty,omitzero"`
	ConsumeUniqueIdAndCounts map[int64]int64 `json:",omitempty,omitzero"`
}
