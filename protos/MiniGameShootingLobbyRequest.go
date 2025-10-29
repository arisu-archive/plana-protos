package protos

type MiniGameShootingLobbyRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
