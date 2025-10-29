package protos

type EventContentStoryStageResultResponse struct {
	ResponsePacket
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	FirstClearReward []ParcelInfo `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
