package protos

type CampaignListResponse struct {
	ResponsePacket
	CampaignChapterClearRewardHistoryDBs []CampaignChapterClearRewardHistoryDB `json:",omitempty,omitzero"`
	StageHistoryDBs                      []CampaignStageHistoryDB              `json:",omitempty,omitzero"`
	StrategyObjecthistoryDBs             []StrategyObjectHistoryDB             `json:",omitempty,omitzero"`
}
