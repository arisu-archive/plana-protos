package protos

type ScenarioSelectResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
