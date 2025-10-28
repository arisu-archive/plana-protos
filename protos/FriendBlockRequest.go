package protos

type FriendBlockRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
