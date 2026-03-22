package protos

type CampaignWithdrawEchelonResponse struct {
	ResponsePacket
	SaveDataDB         CampaignMainStageSaveDB
	WithdrawEchelonDBs []EchelonDB
}
