package protos

type CampaignChapterClearRewardResponse struct {
	ResponsePacket
	CampaignChapterClearRewardHistoryDB CampaignChapterClearRewardHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
