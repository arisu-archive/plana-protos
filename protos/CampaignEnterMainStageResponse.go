package protos

type CampaignEnterMainStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
