package protos

type CampaignMainStageStrategySkipResultResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TacticRank int64 `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs []CharacterDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	FirstClearReward []ParcelInfo `json:",omitempty,omitzero"`
	ThreeStarReward []ParcelInfo `json:",omitempty,omitzero"`
}
