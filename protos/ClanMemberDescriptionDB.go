package protos

type ClanMemberDescriptionDB struct {
	Exp                       int64  `json:",omitempty,omitzero"`
	Comment                   string `json:",omitempty,omitzero"`
	CollectedCharactersCount  int32  `json:",omitempty,omitzero"`
	ArenaSeasonBestRanking    int64  `json:",omitempty,omitzero"`
	ArenaSeasonCurrentRanking int64  `json:",omitempty,omitzero"`
}
