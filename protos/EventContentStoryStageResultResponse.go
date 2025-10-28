package protos

type EventContentStoryStageResultResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	FirstClearReward []ParcelInfo `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
