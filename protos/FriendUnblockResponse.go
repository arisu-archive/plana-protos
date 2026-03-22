package protos

type FriendUnblockResponse struct {
	ResponsePacket
	FriendDBs                []FriendDB
	SentRequestFriendDBs     []FriendDB
	ReceivedRequestFriendDBs []FriendDB
	BlockedUserDBs           []FriendDB
}
