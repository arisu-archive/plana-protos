package protos

type FriendListByIdsResponse struct {
	ResponsePacket
	ListResult []*FriendDB
}
