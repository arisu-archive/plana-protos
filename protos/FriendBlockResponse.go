package protos

type FriendBlockResponse struct {
	ResponsePacket
	FriendDBs                []*FriendDB
	SentRequestFriendDBs     []*FriendDB
	ReceivedRequestFriendDBs []*FriendDB
	BlockedUserDBs           []*FriendDB
}
