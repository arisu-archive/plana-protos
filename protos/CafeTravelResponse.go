package protos

type CafeTravelResponse struct {
	ResponsePacket
	FriendDB  FriendDB
	CafeDBs   []CafeDB
	AllowCopy bool `json:",omitempty,omitzero"`
}
