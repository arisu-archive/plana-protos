package protos

type ClearDeckListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClearDeckDBs []ClearDeckDB `json:",omitempty,omitzero"`
}
