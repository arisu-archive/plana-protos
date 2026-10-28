package protos

type MiniGameTankRaceLobbyRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
