package protos

type CafeApplyTemplateRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TemplateId int64 `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
	UseOtherCafeFurniture bool `json:",omitempty,omitzero"`
}
