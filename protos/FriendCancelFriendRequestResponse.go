package protos

type FriendCancelFriendRequestResponse struct {
	ResponsePacket
	FriendDBs                []*FriendDB
	SentRequestFriendDBs     []*FriendDB
	ReceivedRequestFriendDBs []*FriendDB
	BlockedUserDBs           []*FriendDB
}
