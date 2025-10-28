package protos

type MiniGameCCGGiveupGameResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDB MiniGameCCGSaveDB `json:",omitempty,omitzero"`
}
