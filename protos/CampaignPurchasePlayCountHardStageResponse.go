package protos

type CampaignPurchasePlayCountHardStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
