package protos

type FriendRemoveResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	FriendDBs []FriendDB `json:",omitempty,omitzero"`
	SentRequestFriendDBs []FriendDB `json:",omitempty,omitzero"`
	ReceivedRequestFriendDBs []FriendDB `json:",omitempty,omitzero"`
	BlockedUserDBs []FriendDB `json:",omitempty,omitzero"`
}
