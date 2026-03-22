package protos

type MiniGameCCGReplaceCharacterResponse struct {
	ResponsePacket
	SaveDB         MiniGameCCGSaveDB
	CCGCharacterDB MiniGameCCGCharacterDB
}
