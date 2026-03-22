package protos

type CampaignHealResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB
	SaveDataDB        CampaignMainStageSaveDB
}
