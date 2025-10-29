package protos

type MiniGameCCGGiveupGameResponse struct {
	ResponsePacket
	SaveDB MiniGameCCGSaveDB `json:",omitempty,omitzero"`
}
