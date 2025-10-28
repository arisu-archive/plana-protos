package protos

type EventContentMapMoveResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
	EchelonEntityId int64 `json:",omitempty,omitzero"`
	StrategyObject Strategy `json:",omitempty,omitzero"`
	StrategyObjectParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
}
