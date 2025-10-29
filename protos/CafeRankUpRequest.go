package protos

type CafeRankUpRequest struct {
	RequestPacket
	AccountServerId int64 `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
