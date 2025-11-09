package protos

type CampaignWithdrawEchelonResponse struct {
	ResponsePacket
	SaveDataDB         CampaignMainStageSaveDB `json:",omitempty,omitzero"`
	WithdrawEchelonDBs []EchelonDB             `json:",omitempty,omitzero"`
}
