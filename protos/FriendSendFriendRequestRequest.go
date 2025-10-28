package protos

type FriendSendFriendRequestRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
