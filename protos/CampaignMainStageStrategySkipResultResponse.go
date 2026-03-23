package protos

type CampaignMainStageStrategySkipResultResponse struct {
	ResponsePacket
	TacticRank             int64                   `json:",omitempty,omitzero"`
	CampaignStageHistoryDB *CampaignStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs    []*CharacterDB
	ParcelResultDB         *ParcelResultDB `json:",omitempty,omitzero"`
	FirstClearReward       []*ParcelInfo
	ThreeStarReward        []*ParcelInfo
}
