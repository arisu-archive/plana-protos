package protos

type FriendUnblockRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
