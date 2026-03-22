package protos

type CampaignMapMoveResponse struct {
	ResponsePacket
	SaveDataDB                CampaignMainStageSaveDB
	EchelonEntityId           int64 `json:",omitempty,omitzero"`
	StrategyObject            Strategy
	StrategyObjectParcelInfos []ParcelInfo
}
