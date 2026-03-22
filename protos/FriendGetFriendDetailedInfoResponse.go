package protos

type FriendGetFriendDetailedInfoResponse struct {
	ResponsePacket
	Nickname                       string
	Level                          int64 `json:",omitempty,omitzero"`
	ClanName                       string
	Comment                        string
	FriendCount                    int64 `json:",omitempty,omitzero"`
	FriendCode                     string
	RepresentCharacterUniqueId     int64  `json:",omitempty,omitzero"`
	RepresentCharacterCostumeId    int64  `json:",omitempty,omitzero"`
	CharacterCount                 int64  `json:",omitempty,omitzero"`
	LastNormalCampaignClearStageId *int64 `json:",omitempty,omitzero"`
	LastHardCampaignClearStageId   *int64 `json:",omitempty,omitzero"`
	ArenaRanking                   *int64 `json:",omitempty,omitzero"`
	RaidRanking                    *int64 `json:",omitempty,omitzero"`
	RaidTier                       *int32 `json:",omitempty,omitzero"`
	DetailedAccountInfoDB          DetailedAccountInfoDB
	AttachmentDB                   AccountAttachmentDB
	AssistCharacterDBs             []AssistCharacterDB
}
