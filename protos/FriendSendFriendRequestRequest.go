package protos

type FriendSendFriendRequestRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
