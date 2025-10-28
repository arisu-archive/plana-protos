package protos

type NotificationEventContentReddotRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
