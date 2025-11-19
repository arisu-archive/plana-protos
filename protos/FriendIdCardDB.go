package protos

type FriendIdCardDB struct {
	Level                       int32  `json:",omitempty,omitzero"`
	FriendCode                  string `json:",omitempty,omitzero"`
	Comment                     string `json:",omitempty,omitzero"`
	LastConnectTime             MxTime `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId  int64  `json:",omitempty,omitzero"`
	RepresentCharacterCostumeId int64  `json:",omitempty,omitzero"`
	SearchPermission            bool   `json:",omitempty,omitzero"`
	AutoAcceptFriendRequest     bool   `json:",omitempty,omitzero"`
	CardBackgroundId            int64  `json:",omitempty,omitzero"`
	ShowAccountLevel            bool   `json:",omitempty,omitzero"`
	ShowFriendCode              bool   `json:",omitempty,omitzero"`
	ShowRaidRanking             bool   `json:",omitempty,omitzero"`
	ShowArenaRanking            bool   `json:",omitempty,omitzero"`
	ShowEliminateRaidRanking    bool   `json:",omitempty,omitzero"`
	ArenaRanking                *int64 `json:",omitempty,omitzero"`
	RaidRanking                 *int64 `json:",omitempty,omitzero"`
	RaidTier                    *int32 `json:",omitempty,omitzero"`
	EliminateRaidRanking        *int64 `json:",omitempty,omitzero"`
	EliminateRaidTier           *int32 `json:",omitempty,omitzero"`
	EmblemId                    int64  `json:",omitempty,omitzero"`
}
