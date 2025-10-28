package protos

type FriendGetIdCardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	FriendIdCardDB FriendIdCardDB `json:",omitempty,omitzero"`
}
