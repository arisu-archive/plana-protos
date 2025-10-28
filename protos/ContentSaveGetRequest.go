package protos

type ContentSaveGetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
