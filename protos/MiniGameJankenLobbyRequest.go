package protos

type MiniGameJankenLobbyRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
