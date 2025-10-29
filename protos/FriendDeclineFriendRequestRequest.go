package protos

type FriendDeclineFriendRequestRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
