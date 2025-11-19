package protos

type CampaignEndTurnResponse struct {
	ResponsePacket
	SaveDataDB        CampaignMainStageSaveDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB       `json:",omitempty,omitzero"`
}
