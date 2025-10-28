package protos

type FriendRemoveRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
	Protocol Protocol `json:",omitempty,omitzero"`
}
