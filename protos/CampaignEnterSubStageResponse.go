package protos

type CampaignEnterSubStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB         `json:",omitempty,omitzero"`
	SaveDataDB     CampaignSubStageSaveDB `json:",omitempty,omitzero"`
}
