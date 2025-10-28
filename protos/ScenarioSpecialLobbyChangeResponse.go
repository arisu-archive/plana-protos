package protos

type ScenarioSpecialLobbyChangeResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
