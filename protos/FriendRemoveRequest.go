package protos

type FriendRemoveRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
