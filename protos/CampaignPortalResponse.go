package protos

type CampaignPortalResponse struct {
	ResponsePacket
	CampaignMainStageSaveDB CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
