package protos

type EventContentSubEventLobbyResponse struct {
	ResponsePacket
	EventContentChangeDB EventContentChangeDB `json:",omitempty,omitzero"`
	IsOnSubEvent         bool                 `json:",omitempty,omitzero"`
}
