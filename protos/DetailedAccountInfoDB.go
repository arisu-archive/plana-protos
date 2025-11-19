package protos

type DetailedAccountInfoDB struct {
	AccountId                      int64               `json:",omitempty,omitzero"`
	Nickname                       string              `json:",omitempty,omitzero"`
	Level                          int64               `json:",omitempty,omitzero"`
	ClanName                       string              `json:",omitempty,omitzero"`
	Comment                        string              `json:",omitempty,omitzero"`
	FriendCount                    int64               `json:",omitempty,omitzero"`
	FriendCode                     string              `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId     int64               `json:",omitempty,omitzero"`
	CharacterCount                 int64               `json:",omitempty,omitzero"`
	LastNormalCampaignClearStageId *int64              `json:",omitempty,omitzero"`
	LastHardCampaignClearStageId   *int64              `json:",omitempty,omitzero"`
	ArenaRanking                   *int64              `json:",omitempty,omitzero"`
	RaidRanking                    *int64              `json:",omitempty,omitzero"`
	RaidTier                       *int32              `json:",omitempty,omitzero"`
	EliminateRaidRanking           *int64              `json:",omitempty,omitzero"`
	EliminateRaidTier              *int32              `json:",omitempty,omitzero"`
	AssistCharacterDBs             []AssistCharacterDB `json:",omitempty,omitzero"`
}
