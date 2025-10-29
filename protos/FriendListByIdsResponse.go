package protos

type FriendListByIdsResponse struct {
	ResponsePacket
	ListResult []FriendDB `json:",omitempty,omitzero"`
}
