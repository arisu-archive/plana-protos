package protos

type ScenarioEndTurnRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
