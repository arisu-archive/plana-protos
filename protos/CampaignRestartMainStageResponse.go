package protos

type CampaignRestartMainStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
