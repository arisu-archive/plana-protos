package protos

type ScenarioLobbyStudentChangeRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	LobbyStudents []int64 `json:",omitempty,omitzero"`
	LobbyStudentsBefore []int64 `json:",omitempty,omitzero"`
}
