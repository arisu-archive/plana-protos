package protos

type CampaignHealResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB       `json:",omitempty,omitzero"`
	SaveDataDB        CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
