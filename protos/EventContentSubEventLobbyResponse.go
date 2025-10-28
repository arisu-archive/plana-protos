package protos

type EventContentSubEventLobbyResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentChangeDB EventContentChangeDB `json:",omitempty,omitzero"`
	IsOnSubEvent bool `json:",omitempty,omitzero"`
}
