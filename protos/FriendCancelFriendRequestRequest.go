package protos

type FriendCancelFriendRequestRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
