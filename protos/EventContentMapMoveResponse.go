package protos

type EventContentMapMoveResponse struct {
	ResponsePacket
	SaveDataDB                EventContentMainStageSaveDB
	EchelonEntityId           int64 `json:",omitempty,omitzero"`
	StrategyObject            Strategy
	StrategyObjectParcelInfos []ParcelInfo
}
