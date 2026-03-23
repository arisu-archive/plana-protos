package protos

type FriendSearchResponse struct {
	ResponsePacket
	SearchResult []*FriendDB
}
