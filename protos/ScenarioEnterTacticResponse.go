package protos

type ScenarioEnterTacticResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
