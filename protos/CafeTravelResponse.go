package protos

type CafeTravelResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	FriendDB FriendDB `json:",omitempty,omitzero"`
	CafeDBs []CafeDB `json:",omitempty,omitzero"`
}
