package protos

type ScenarioEndTurnResponse struct {
	ResponsePacket
	SaveDataDB        StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB        `json:",omitempty,omitzero"`
	ScenarioIds       []int64                  `json:",omitempty,omitzero"`
}
