package protos

type FriendCheckRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
