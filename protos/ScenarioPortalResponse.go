package protos

type ScenarioPortalResponse struct {
	ResponsePacket
	StoryStrategyStageSaveDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	ScenarioIds              []int64                  `json:",omitempty,omitzero"`
}
