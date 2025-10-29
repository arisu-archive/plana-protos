package protos

type ScenarioRetreatRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
