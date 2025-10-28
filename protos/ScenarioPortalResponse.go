package protos

type ScenarioPortalResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StoryStrategyStageSaveDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	ScenarioIds []int64 `json:",omitempty,omitzero"`
}
