package protos

type ScenarioLobbyStudentChangeRequest struct {
	RequestPacket
	LobbyStudents       []int64
	LobbyStudentsBefore []int64
}
