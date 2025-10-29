package protos

type MemoryLobbyDB struct {
	ParcelBase
	MemoryLobbyUniqueId int64 `json:",omitempty,omitzero"`
}
