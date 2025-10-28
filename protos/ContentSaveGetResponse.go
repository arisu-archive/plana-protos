package protos

type ContentSaveGetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	HasValidData bool `json:",omitempty,omitzero"`
	ContentSaveDB ContentSaveDB `json:",omitempty,omitzero"`
	EventContentChangeDB EventContentChangeDB `json:",omitempty,omitzero"`
}
