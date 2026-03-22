package protos

type DetailedAccountInfoDB struct {
	AccountId                       int64 `json:",omitempty,omitzero"`
	Nickname                        string
	Level                           int64 `json:",omitempty,omitzero"`
	ClanName                        string
	Comment                         string
	FriendCount                     int64 `json:",omitempty,omitzero"`
	FriendCode                      string
	RepresentCharacterUniqueId      int64  `json:",omitempty,omitzero"`
	CharacterCount                  int64  `json:",omitempty,omitzero"`
	LastNormalCampaignClearStageId  *int64 `json:",omitempty,omitzero"`
	LastHardCampaignClearStageId    *int64 `json:",omitempty,omitzero"`
	ArenaRanking                    *int64 `json:",omitempty,omitzero"`
	RaidRanking                     *int64 `json:",omitempty,omitzero"`
	RaidTier                        *int32 `json:",omitempty,omitzero"`
	EliminateRaidRanking            *int64 `json:",omitempty,omitzero"`
	EliminateRaidTier               *int32 `json:",omitempty,omitzero"`
	MultiFloorRaidClearedDifficulty *int64 `json:",omitempty,omitzero"`
	AssistCharacterDBs              []AssistCharacterDB
}
