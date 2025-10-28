package protos

type FriendListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
