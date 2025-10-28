package protos

type CampaignPortalResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CampaignMainStageSaveDB CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
