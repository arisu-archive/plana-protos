package protos

type CampaignEnterSubStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB CampaignSubStageSaveDB `json:",omitempty,omitzero"`
}
