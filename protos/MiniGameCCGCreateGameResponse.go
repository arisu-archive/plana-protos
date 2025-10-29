package protos

type MiniGameCCGCreateGameResponse struct {
	ResponsePacket
	CCGSaveDB MiniGameCCGSaveDB `json:",omitempty,omitzero"`
}
