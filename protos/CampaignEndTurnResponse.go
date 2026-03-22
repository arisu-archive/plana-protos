package protos

type CampaignEndTurnResponse struct {
	ResponsePacket
	SaveDataDB        CampaignMainStageSaveDB
	AccountCurrencyDB AccountCurrencyDB
}
