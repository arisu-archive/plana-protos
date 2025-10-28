package protos

type EventContentPurchasePlayCountHardStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
