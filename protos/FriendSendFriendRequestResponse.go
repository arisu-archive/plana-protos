package protos

type FriendSendFriendRequestResponse struct {
	ResponsePacket
	FriendDBs                []FriendDB `json:",omitempty,omitzero"`
	SentRequestFriendDBs     []FriendDB `json:",omitempty,omitzero"`
	ReceivedRequestFriendDBs []FriendDB `json:",omitempty,omitzero"`
	BlockedUserDBs           []FriendDB `json:",omitempty,omitzero"`
}
