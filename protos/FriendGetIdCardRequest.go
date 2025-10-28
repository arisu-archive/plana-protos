package protos

type FriendGetIdCardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
