package protos

type FriendCheckResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
