package protos

type CafeApplyTemplateRequest struct {
	RequestPacket
	TemplateId            int64 `json:",omitempty,omitzero"`
	CafeDBId              int64 `json:",omitempty,omitzero"`
	UseOtherCafeFurniture bool  `json:",omitempty,omitzero"`
}
