package protos

type ScenarioConfirmMainStageRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
