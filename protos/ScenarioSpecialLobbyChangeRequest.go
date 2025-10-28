package protos

type ScenarioSpecialLobbyChangeRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MemoryLobbyId int64 `json:",omitempty,omitzero"`
	MemoryLobbyIdBefore int64 `json:",omitempty,omitzero"`
}
