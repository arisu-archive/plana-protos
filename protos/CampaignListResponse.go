package protos

type CampaignListResponse struct {
	ResponsePacket
	CampaignChapterClearRewardHistoryDBs []CampaignChapterClearRewardHistoryDB
	StageHistoryDBs                      []CampaignStageHistoryDB
	StrategyObjecthistoryDBs             []StrategyObjectHistoryDB
}
