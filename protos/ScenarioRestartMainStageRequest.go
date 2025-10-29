package protos

type ScenarioRestartMainStageRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
