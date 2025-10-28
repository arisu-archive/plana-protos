package protos

type ScenarioListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
