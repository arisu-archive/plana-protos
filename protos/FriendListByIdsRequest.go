package protos

type FriendListByIdsRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetAccountIds []int64 `json:",omitempty,omitzero"`
}
