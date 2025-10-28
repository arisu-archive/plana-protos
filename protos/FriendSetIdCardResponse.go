package protos

type FriendSetIdCardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
