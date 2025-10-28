package protos

type FriendListResponse struct {
	ResponsePacket
	FriendIdCardDB FriendIdCardDB `json:",omitempty,omitzero"`
	Protocol Protocol `json:",omitempty,omitzero"`
	IdCardBackgroundDBs []IdCardBackgroundDB `json:",omitempty,omitzero"`
	FriendDBs []FriendDB `json:",omitempty,omitzero"`
	SentRequestFriendDBs []FriendDB `json:",omitempty,omitzero"`
	ReceivedRequestFriendDBs []FriendDB `json:",omitempty,omitzero"`
	BlockedUserDBs []FriendDB `json:",omitempty,omitzero"`
}
