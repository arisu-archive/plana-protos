package protos

type CampaignListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CampaignChapterClearRewardHistoryDBs []CampaignChapterClearRewardHistoryDB `json:",omitempty,omitzero"`
	StageHistoryDBs []CampaignStageHistoryDB `json:",omitempty,omitzero"`
	StrategyObjecthistoryDBs []StrategyObjectHistoryDB `json:",omitempty,omitzero"`
}
