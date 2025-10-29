package protos

type FriendGetIdCardResponse struct {
	ResponsePacket
	FriendIdCardDB FriendIdCardDB `json:",omitempty,omitzero"`
}
