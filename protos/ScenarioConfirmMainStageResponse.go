package protos

type ScenarioConfirmMainStageResponse struct {
	ResponsePacket
	ParcelResultDB *ParcelResultDB           `json:",omitempty,omitzero"`
	SaveDataDB     *StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	ScenarioIds    []int64
}
