package protos

type FriendListByIdsRequest struct {
	RequestPacket
	TargetAccountIds []int64
}
