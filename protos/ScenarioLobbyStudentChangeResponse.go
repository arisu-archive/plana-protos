package protos

type ScenarioLobbyStudentChangeResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
