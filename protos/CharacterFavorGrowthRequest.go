package protos

type CharacterFavorGrowthRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetCharacterDBId int64 `json:",omitempty,omitzero"`
	ConsumeItemDBIdsAndCounts map[int64]int32 `json:",omitempty,omitzero"`
}
