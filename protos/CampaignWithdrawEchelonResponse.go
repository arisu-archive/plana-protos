package protos

type CampaignWithdrawEchelonResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB CampaignMainStageSaveDB `json:",omitempty,omitzero"`
	WithdrawEchelonDBs []EchelonDB `json:",omitempty,omitzero"`
}
