package protos

type NotificationLobbyCheckRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
