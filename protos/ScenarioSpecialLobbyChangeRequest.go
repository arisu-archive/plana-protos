package protos

type ScenarioSpecialLobbyChangeRequest struct {
	RequestPacket
	MemoryLobbyId       int64 `json:",omitempty,omitzero"`
	MemoryLobbyIdBefore int64 `json:",omitempty,omitzero"`
}
