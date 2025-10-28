package protos

type ScenarioSkipResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
