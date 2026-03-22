package protos

type EventContentSubEventLobbyResponse struct {
	ResponsePacket
	EventContentChangeDB EventContentChangeDB
	IsOnSubEvent         bool `json:",omitempty,omitzero"`
}
