package protos

type FriendListByIdsRequest struct {
	RequestPacket
	TargetAccountIds []int64 `json:",omitempty,omitzero"`
}
