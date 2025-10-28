package protos

type ScenarioConfirmMainStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	ScenarioIds []int64 `json:",omitempty,omitzero"`
}
