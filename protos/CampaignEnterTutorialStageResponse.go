package protos

type CampaignEnterTutorialStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	SaveDataDB     CampaignTutorialStageSaveDB
}
