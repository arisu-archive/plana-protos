package protos

type EventContentMapMoveResponse struct {
	ResponsePacket
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
	EchelonEntityId int64 `json:",omitempty,omitzero"`
	StrategyObject Strategy `json:",omitempty,omitzero"`
	StrategyObjectParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
}
