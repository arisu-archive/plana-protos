package protos

type ScenarioEndTurnResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ScenarioIds []int64 `json:",omitempty,omitzero"`
}
