package protos

type FriendGetFriendDetailedInfoResponse struct {
	ResponsePacket
	Nickname string `json:",omitempty,omitzero"`
	Level int64 `json:",omitempty,omitzero"`
	ClanName string `json:",omitempty,omitzero"`
	Comment string `json:",omitempty,omitzero"`
	FriendCount int64 `json:",omitempty,omitzero"`
	FriendCode string `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId int64 `json:",omitempty,omitzero"`
	RepresentCharacterCostumeId int64 `json:",omitempty,omitzero"`
	CharacterCount int64 `json:",omitempty,omitzero"`
	LastNormalCampaignClearStageId *int64 `json:",omitempty,omitzero"`
	LastHardCampaignClearStageId *int64 `json:",omitempty,omitzero"`
	ArenaRanking *int64 `json:",omitempty,omitzero"`
	RaidRanking *int64 `json:",omitempty,omitzero"`
	RaidTier *int32 `json:",omitempty,omitzero"`
	DetailedAccountInfoDB DetailedAccountInfoDB `json:",omitempty,omitzero"`
	AttachmentDB AccountAttachmentDB `json:",omitempty,omitzero"`
	AssistCharacterDBs []AssistCharacterDB `json:",omitempty,omitzero"`
}
