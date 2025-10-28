package protos

type MiniGameCCGCreateGameResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CCGSaveDB MiniGameCCGSaveDB `json:",omitempty,omitzero"`
}
