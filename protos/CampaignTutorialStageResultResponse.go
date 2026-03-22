package protos

type CampaignTutorialStageResultResponse struct {
	ResponsePacket
	CampaignStageHistoryDB CampaignStageHistoryDB
	ParcelResultDB         ParcelResultDB
	ClearReward            []ParcelInfo
	FirstClearReward       []ParcelInfo
}
