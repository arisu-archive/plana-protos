package protos

type ContentSaveGetResponse struct {
	ResponsePacket
	HasValidData         bool                 `json:",omitempty,omitzero"`
	ContentSaveDB        ContentSaveDB        `json:",omitempty,omitzero"`
	EventContentChangeDB EventContentChangeDB `json:",omitempty,omitzero"`
}
