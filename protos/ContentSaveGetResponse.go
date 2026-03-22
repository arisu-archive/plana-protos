package protos

type ContentSaveGetResponse struct {
	ResponsePacket
	HasValidData         bool `json:",omitempty,omitzero"`
	ContentSaveDB        ContentSaveDB
	EventContentChangeDB EventContentChangeDB
}
