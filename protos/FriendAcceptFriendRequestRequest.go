package protos

type FriendAcceptFriendRequestRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
