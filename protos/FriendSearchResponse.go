package protos

type FriendSearchResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SearchResult []FriendDB `json:",omitempty,omitzero"`
}
