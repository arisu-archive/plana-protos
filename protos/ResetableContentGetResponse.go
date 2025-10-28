package protos

type ResetableContentGetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ResetableContentValueDBs []ResetableContentValueDB `json:",omitempty,omitzero"`
}
