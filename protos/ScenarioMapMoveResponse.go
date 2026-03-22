package protos

type ScenarioMapMoveResponse struct {
	ResponsePacket
	SaveDataDB                StoryStrategyStageSaveDB
	ScenarioIds               []int64
	EchelonEntityId           int64 `json:",omitempty,omitzero"`
	StrategyObject            Strategy
	StrategyObjectParcelInfos []ParcelInfo
}
