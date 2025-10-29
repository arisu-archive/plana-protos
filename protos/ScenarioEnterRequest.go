package protos

type ScenarioEnterRequest struct {
	RequestPacket
	ScenarioId int64 `json:",omitempty,omitzero"`
}
