package protos

type MiniGameCCGLobbyRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
