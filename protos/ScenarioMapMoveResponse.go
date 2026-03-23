package protos

type ScenarioMapMoveResponse struct {
	ResponsePacket
	SaveDataDB                *StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	ScenarioIds               []int64
	EchelonEntityId           int64     `json:",omitempty,omitzero"`
	StrategyObject            *Strategy `json:",omitempty,omitzero"`
	StrategyObjectParcelInfos []ParcelInfo
}
