package protos

type CafeTravelResponse struct {
	ResponsePacket
	FriendDB  *FriendDB `json:",omitempty,omitzero"`
	CafeDBs   []CafeDB
	AllowCopy bool `json:",omitempty,omitzero"`
}
