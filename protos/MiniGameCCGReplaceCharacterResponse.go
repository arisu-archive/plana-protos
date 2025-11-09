package protos

type MiniGameCCGReplaceCharacterResponse struct {
	ResponsePacket
	SaveDB         MiniGameCCGSaveDB      `json:",omitempty,omitzero"`
	CCGCharacterDB MiniGameCCGCharacterDB `json:",omitempty,omitzero"`
}
