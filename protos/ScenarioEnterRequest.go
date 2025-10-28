package protos

type ScenarioEnterRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScenarioId int64 `json:",omitempty,omitzero"`
}
