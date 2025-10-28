package protos

type CampaignChapterClearRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CampaignChapterClearRewardHistoryDB CampaignChapterClearRewardHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
