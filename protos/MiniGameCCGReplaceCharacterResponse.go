package protos

type MiniGameCCGReplaceCharacterResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDB MiniGameCCGSaveDB `json:",omitempty,omitzero"`
	CCGCharacterDB MiniGameCCGCharacterDB `json:",omitempty,omitzero"`
}
