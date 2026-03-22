package protos

type FriendListResponse struct {
	ResponsePacket
	FriendIdCardDB           FriendIdCardDB
	IdCardBackgroundDBs      []IdCardBackgroundDB
	FriendDBs                []FriendDB
	SentRequestFriendDBs     []FriendDB
	ReceivedRequestFriendDBs []FriendDB
	BlockedUserDBs           []FriendDB
}
