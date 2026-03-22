package protos

type CampaignConfirmMainStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	SaveDataDB     CampaignMainStageSaveDB
}
