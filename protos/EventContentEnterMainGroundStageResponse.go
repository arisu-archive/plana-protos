package protos

type EventContentEnterMainGroundStageResponse struct {
	ResponsePacket
	ParcelResultDB         ParcelResultDB
	SaveDataDB             EventContentMainGroundStageSaveDB
	CampaignStageHistoryDB CampaignStageHistoryDB
}
