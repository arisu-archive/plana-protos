package protos

type ScenarioLobbyStudentChangeRequest struct {
	RequestPacket
	LobbyStudents       []int64 `json:",omitempty,omitzero"`
	LobbyStudentsBefore []int64 `json:",omitempty,omitzero"`
}
