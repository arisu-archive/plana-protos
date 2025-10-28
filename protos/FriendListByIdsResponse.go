package protos

type FriendListByIdsResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ListResult []FriendDB `json:",omitempty,omitzero"`
}
