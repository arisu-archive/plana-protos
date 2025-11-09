package protos

type CampaignMapMoveResponse struct {
	ResponsePacket
	SaveDataDB                CampaignMainStageSaveDB `json:",omitempty,omitzero"`
	EchelonEntityId           int64                   `json:",omitempty,omitzero"`
	StrategyObject            Strategy                `json:",omitempty,omitzero"`
	StrategyObjectParcelInfos []ParcelInfo            `json:",omitempty,omitzero"`
}
