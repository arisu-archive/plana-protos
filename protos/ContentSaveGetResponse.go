package protos

type ContentSaveGetResponse struct {
	ResponsePacket
	HasValidData         bool `json:",omitempty,omitzero"`
	EventContentChangeDB EventContentChangeDB
}
