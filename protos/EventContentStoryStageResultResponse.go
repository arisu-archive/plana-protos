package protos

type EventContentStoryStageResultResponse struct {
	ResponsePacket
	CampaignStageHistoryDB    CampaignStageHistoryDB
	ParcelResultDB            ParcelResultDB
	FirstClearReward          []ParcelInfo
	EventContentCollectionDBs []EventContentCollectionDB
}
