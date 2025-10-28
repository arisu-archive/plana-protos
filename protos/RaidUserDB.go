package protos

type RaidUserDB struct {
	AccountId int64 `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId int64 `json:",omitempty,omitzero"`
	RepresentCharacterCostumeId int64 `json:",omitempty,omitzero"`
	Level int64 `json:",omitempty,omitzero"`
	Nickname string `json:",omitempty,omitzero"`
	Tier int32 `json:",omitempty,omitzero"`
	Rank int64 `json:",omitempty,omitzero"`
	BestRankingPoint int64 `json:",omitempty,omitzero"`
	BestRankingPointDetail float64 `json:",omitempty,omitzero"`
	AccountAttachmentDB AccountAttachmentDB `json:",omitempty,omitzero"`
}
