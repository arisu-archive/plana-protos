package protos

type EventContentStoryStageResultResponse struct {
	ResponsePacket
	CampaignStageHistoryDB    *CampaignStageHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB            *ParcelResultDB         `json:",omitempty,omitzero"`
	FirstClearReward          []ParcelInfo
	EventContentCollectionDBs []EventContentCollectionDB
}
