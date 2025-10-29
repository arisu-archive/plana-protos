package protos

type EventContentPurchasePlayCountHardStageResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
