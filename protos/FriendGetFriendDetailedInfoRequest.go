package protos

type FriendGetFriendDetailedInfoRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	FriendAccountId int64 `json:",omitempty,omitzero"`
}
