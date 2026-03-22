package protos

type EventContentPurchasePlayCountHardStageResponse struct {
	ResponsePacket
	AccountCurrencyDB      AccountCurrencyDB
	CampaignStageHistoryDB CampaignStageHistoryDB
}
