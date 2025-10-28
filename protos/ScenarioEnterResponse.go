package protos

type ScenarioEnterResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
