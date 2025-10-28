package protos

type CafeRankUpRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountServerId int64 `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
