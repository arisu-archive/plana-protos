package protos

type MemoryLobbyListResponse struct {
	ResponsePacket
	MemoryLobbyDBs []MemoryLobbyDB
}
