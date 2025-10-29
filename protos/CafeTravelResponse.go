package protos

type CafeTravelResponse struct {
	ResponsePacket
	FriendDB FriendDB `json:",omitempty,omitzero"`
	CafeDBs []CafeDB `json:",omitempty,omitzero"`
}
