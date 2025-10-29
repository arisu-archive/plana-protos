package protos

type FriendBlockRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
