package protos

type FriendListResponse struct {
	ResponsePacket
	FriendIdCardDB           *FriendIdCardDB `json:",omitempty,omitzero"`
	IdCardBackgroundDBs      []*IdCardBackgroundDB
	FriendDBs                []*FriendDB
	SentRequestFriendDBs     []*FriendDB
	ReceivedRequestFriendDBs []*FriendDB
	BlockedUserDBs           []*FriendDB
}
