package protos

type FriendGetFriendDetailedInfoRequest struct {
	RequestPacket
	FriendAccountId int64 `json:",omitempty,omitzero"`
}
